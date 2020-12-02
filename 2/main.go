package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type policy struct {
	LowerValue int
	UpperValue int
	Character  string
}

type line struct {
    Pol policy
    Password string
}


//1-3 a: abcde
//1-3 b: cdefg
//2-9 c: ccccccccc

func parsePolicy(input string) policy {
    pol := policy{}
    splitOnSpace := strings.Split(input, " ")
    pol.Character = splitOnSpace[1]
    splitRange := strings.Split(splitOnSpace[0], "-")
    pol.LowerValue, _ = strconv.Atoi(splitRange[0])
    pol.UpperValue, _ = strconv.Atoi(splitRange[1])
    return pol
}

func parseLine(input string) line {
    l := line{}
    spl := strings.Split(input, ": ")
    l.Pol = parsePolicy(spl[0])
    l.Password = spl[1]
    return l
}

func (l *line) valid() bool {
    return l.Pol.LowerValue <= strings.Count(l.Password, l.Pol.Character) && 
           l.Pol.UpperValue >= strings.Count(l.Password, l.Pol.Character)
}

func (l *line) valid2() bool {
    return (string(l.Password[l.Pol.LowerValue - 1]) == l.Pol.Character && string(l.Password[l.Pol.UpperValue - 1]) != l.Pol.Character) ||
           (string(l.Password[l.Pol.LowerValue - 1]) != l.Pol.Character && string(l.Password[l.Pol.UpperValue - 1]) == l.Pol.Character)
}

func readInput(file string) []string {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	text := string(dat)
	return strings.Split(text, "\n")
}

func main () {
    inputData := readInput("input")
    count := 0
    for _, l := range inputData {
        pl := parseLine(l)
        if pl.valid2() {
            count++
        }
    }
    fmt.Println(count)
}