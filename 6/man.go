package main

import (
    "strings"
    "io/ioutil"
    "fmt"
    "log"
)

func readInput(file string) []string {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	text := string(dat)
	return strings.Split(text, "\n\n")
}

func Union(m1, m2 map[byte]bool) map[byte]bool {
    mOut := make(map[byte]bool, len(m1))
    for b1 := range m1 {
        mOut[b1] = true
    }
    for b2 := range m2 {
        mOut[b2] = true
    }
    return mOut
}

func Intersect(m1, m2 map[byte]bool) map[byte]bool {
    mOut := make(map[byte]bool, len(m1))
    for b1 := range m1 {
        _, exists := m2[b1]
        if exists {
            mOut[b1] = true
        }
    }
    return mOut
}

func main() {
    groupedLines := readInput("input")
    
    totalCount := 0
    for _, group := range groupedLines {
        var groupSet map[byte]bool
        
        personList := strings.Split(group, "\n")
        for _, person := range(personList) {
            personSet := make(map[byte]bool, 0)
            for i := 0; i < len(person); i++ {
                personSet[person[i]] = true
            }
            if groupSet == nil {
                groupSet = personSet
            } else {
                groupSet = Intersect(groupSet, personSet)
            }
        }
        groupCount := len(groupSet)
        fmt.Printf("group count: %d\n", groupCount)
        totalCount += groupCount
    }
    //4574 too high
    fmt.Printf("total size %d\n", totalCount)
}