package main

import (
    "io/ioutil"
    "log"
    "fmt"
    "strings"
)

type BoardingPass = string
type Seat struct {
    Row, Col int
}

func Power(x, n int) int{
    if n == 0 {
        return 1
    }
    return x * Power(x, n - 1)
}

func PassToInt(bp BoardingPass) Seat {
    s := Seat{0,0}
    for i := 0; i < 7; i++ {
        if bp[i] == 'B' {
            s.Row += Power(2, 6 - i)
        } else if bp[i] != 'F' {
            return Seat{}
        }
    }
    for j := 7; j < 10; j++ {
        if bp[j] == 'R' {
            s.Col += Power(2, 9 - j)
        } else if bp[j] != 'L' {
            return Seat{}
        }
    }
    return s
}

func (s Seat) getID() int {
    return s.Row * 8 + s.Col
}

func main() {
    passes := readInput("input")
    maxID := 0
    idList := make(map[int]BoardingPass)
    for _, p := range passes {
        curID := PassToInt(p).getID()
        idList[curID] = p
        if curID > maxID {
            maxID = curID
        }
    }
    
    fmt.Printf("maximum id: %d\n", maxID)
    for i := 1; i < maxID; i++ {
        _, exists := idList[i]
        if(!exists) {
            fmt.Printf("%d not present\n", i)
        }
    }
}

func readInput(file string) []string {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	text := string(dat)
	return strings.Split(text, "\n")
}
