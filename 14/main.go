package main

import (
    "io/ioutil"
    "strings"
    "strconv"
    "log"
    "math"
)

func readInput(file string) []string {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	text := string(dat)
	return strings.Split(text, "\n")
}

type MemoryInstruct struct {
    Address int
    Value   int
}

type Bitmask string

type Bitstring string

func Reverse(s Bitstring) Bitstring {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return Bitstring(string(runes))
}

func iToBitstring(n float64) Bitstring {
    s := ""
    i := float64(35)
    for i >= 0 {
        if n >= math.Exp2(i) {
            s += "1"
            n -= math.Exp2(i)
        } else {
            s += "0"
        }
        i--
    }
    return Bitstring(s)
}

func BitstringToFloat64(bs Bitstring) float64 {
    br := Reverse(bs)
    result := float64(0)
    for i, b := range br {
        bit, _ := strconv.Atoi(string(b))
        result += float64(bit) * math.Exp2(float64(i))
    }
    return result
}

func maskString(v Bitstring, m Bitmask) Bitstring {
    result := ""
    for i, b := range m {
        if b == 'X' {
            result += string(v[i])
        } else {
            result += string(b)
        }
    }
    return Bitstring(result)
}

func maskAddress(v Bitstring, m Bitmask) []*Bitstring {
    result := make([]*Bitstring, 1)
    var emptyString = Bitstring("")
    result[0] = &emptyString
    
    for i, b := range m {
        switch b {
        case 'X': // floating
            for _, r := range result {
                newRzero := *r + "0"
                newRone := *r + "1"
                r = &newRzero
                result = append(result, &newRone)
            }
        case '0': //unchanged from v
            for _, r := range result {
                newR := *r + Bitstring(v[i])
                r = &newR
            }
        case '1': // output 1
            for _, r := range result {
                newR := *r + "1"
                r = &newR
            }
        }
    }
    log.Printf("address was %s\n", v)
    log.Printf("mask    was %s\n", m)
    for _, a := range result {
        log.Printf("%s\n", *a)
    }
    return result
}

func main() {
    cs := readInput("input")
    var mask Bitmask = ""
    var memory map[int]float64
    memory = make(map[int]float64, 0)
    for _, c := range cs {
        if strings.HasPrefix(c, "mask") {
            m := strings.TrimPrefix(c, "mask = ")
            mask = Bitmask(m)
            log.Printf("new mask %s\n", mask)
        } else if strings.HasPrefix(c, "mem") {
            ss := strings.Split(c, " = ")
            newValue, _ := strconv.Atoi(ss[1])
            address, _ := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(ss[0], "mem["), "]"))
            allMaskedAddress := maskAddress(iToBitstring(float64(address)), mask)
            for _, a := range allMaskedAddress {
                add := int(BitstringToFloat64(*a))
                memory[add] = float64(newValue)
                log.Printf("updated [%d] to %d\n", add, memory[add])
            }
            
            
        }
    }
    var result = float64(0)
    for _, v := range memory {
        result += v
    }
    log.Printf("sum of memory %d\n", int(result))
}