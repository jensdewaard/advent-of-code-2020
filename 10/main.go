package main

import (
    "strings"
    "strconv"
    "log"
    "io/ioutil"
    "sort"
)

func readInput(file string) []string {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	text := string(dat)
	return strings.Split(text, "\n")
}

type fn func(string) (int, error)

func mapStringSlice(ss []string, f fn) []int {
    is := make([]int, 0)
    for _, s := range ss {
        i, e := f(s)
        if e != nil {
            log.Fatal(e)
        }
        is = append(is, i)
    }
    return is
}

func calculateDeltas(adapters []int) {
    ds := make([]int, 0)
    for i := range adapters {
        if i == 0 {
            continue
        }
        d := adapters[i] - adapters[i-1]
        ds = append(ds, d)
    }
    log.Printf("%v", ds)
    occ3 := countOccurence(ds, 3)
    occ1 := countOccurence(ds, 1)
    log.Printf("3: %d", occ3)
    log.Printf("1: %d", occ1)
    log.Printf("1 * 3: %d", occ3 * occ1)
}

func main() {
    adapters := mapStringSlice(readInput("input"), strconv.Atoi)
    adapters = append(adapters, 0, max(adapters) + 3)
    sort.Ints(adapters)
    
    //calculateDeltas(adapters)
    memoization = make(map[int]int, 0)
    log.Printf("arrangements possible %d", arrangements(adapters))
}

var memoization map[int]int

func arrangements(is []int) int {
    if len(is) == 0 {
        return 0
    }
    if len(is) == 1 {
        return 1
    }
    if len(is) == 2 {
        if is[0] + 3 >= is[1] {
            return 1
        }
        return 0
    }
    branches := 0
    curr := is[0]

    result, exists := memoization[curr]
    if exists {
        return result
    }
    
    if is[1] <= curr + 3 {
        //we can hop to the next (this should always be true)
        branches++
    }
    if len(is) > 2 && is[2] <= curr + 3 {
        // we skip one
        branches++
    }
    if len(is) > 3 && is[3] <= curr + 3 {
        //we can skip two
        branches++
    }
    switch branches {
    case 1:
        // the other adapters are out of reach, so we pick the next one
        result = arrangements(is[1:])
        memoization[curr] = result
        return result
    case 2:
        // the adapter at place 1 is optional, so we can double the 
        // number of ways the next adapters can be arranged
        result = arrangements(is[1:]) + arrangements(is[2:])
        memoization[curr] = result
        return result
    case 3:
        // we can choose how much adapters to skip
        result = arrangements(is[1:]) + arrangements(is[2:]) + arrangements(is[3:])
        memoization[curr] = result
        return result
    }
    return 1
}

func countOccurence(is []int, t int) int {
    c := 0
    for _, i := range is {
        if i == t {
            c++
        }
    }
    return c
}

func max(is []int) int {
    m := 0
    for _, i := range is {
        if i > m {
            m = i
        }
    }
    return m
}