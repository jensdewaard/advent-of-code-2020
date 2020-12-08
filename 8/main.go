package main

import (
    "errors"
    "fmt"
    "strings"
    "strconv"
    "io/ioutil"
    "log"
)

var accumulator = 0
var runInstructionsIndices = make([]int, 0)
var ErrInfiniteLoop = errors.New("infinite loop")

type instruction struct {
    code string
    value int
    run bool
}

type program []instruction

func (i *instruction) String() {
    fmt.Sprintf("%s %d", i.code, i.value)
}

func parseInput(s string) instruction {
    ss := strings.Split(s, " ")
    com := ss[0]
    value, _ := strconv.Atoi(ss[1])
    return instruction{com, value, false}
}

func Run(p program, ctr, ptr int) (int, error) {
    if ptr >= len(p) {
        return accumulator, nil
    }
    c := &p[ptr]
    if c.run {
        return 0, ErrInfiniteLoop
    }
    runInstructionsIndices = append(runInstructionsIndices, ptr)
    switch c.code {
    case "acc":
        accumulator += c.value
        ptr++
    case "jmp":
        ptr += c.value
    case "nop":
        ptr++
    }
    c.run = true
    
    //fmt.Printf("[%d] %d %s %d\n", ctr, ptr, c.code, c.value)
    return Run(p, ctr + 1, ptr)
}

func readInput(file string) []string {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	text := string(dat)
	return strings.Split(text, "\n")
}

func main() {
    ls := readInput("input")
    p := make(program, 0)
    for _, l := range ls {
        p = append(p, parseInput(l))
    }
    var p2 = make(program, len(p))
    copy(p2, p)
    // initial run
    Run(p2, 0,0) // we know it failed
    
    candidates := make([]int, 0)
    for _, i := range runInstructionsIndices {
        if p[i].code == "nop" || p[i].code == "jmp" {
            candidates = append(candidates, i)
        }
    }
    // try changing the program on each candidate
    for _, i := range candidates {
        fmt.Printf("attempt changing %d\n", i)
        accumulator = 0
        _, err := Run(fixProgram(p, i), 0, 0)
        if err == nil {
            fmt.Printf("acc %d\n", accumulator)
            break
        }
        
    }
    

    //fmt.Printf("acc %d %v\n", r, err)
    
}

func listCandidates(p []instruction) {
    l := len(p)
    for i, c := range p {
        target := i + c.value
        switch c.code {
        case "acc":
            continue
        case "nop":
            if target >= 625 && target <= l {
                fmt.Printf("%d: %s %d\n", i, c.code, c.value)
            }
        case "jmp":
            if i + c.value >= 625 && target <= l {
                fmt.Printf("%d: %s %d\n", i, c.code, c.value)
            }
        }
    }
}

// the fix is done by changing a single NOP to a JMP 
// or a single JMP to a NOP. Therefore, somewhere in the program
// there either already exists a JMP to len(program) that can be
// reached by changing a JMP to a NOP, or there is a NOP +x that
// would jump to the len(program) if it were a JMP
func fixProgram(p program, i int) program {
    var p2 = make(program, len(p))
    copy(p2, p)
    switch p2[i].code {
    case "nop":
        p2[i].code = "jmp"
    case "jmp":
        p2[i].code = "nop"
    }
    return p2
}