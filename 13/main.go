package main

import(
    "io/ioutil"
    "log"
    "strings"
    "strconv"
    "math"
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

type fn func (string) (int, error)

func mapStringSlice(ss []string, f fn) []int {
    as := make([]int, 0)
    for _, s := range ss {
        if s == "x" {
            continue;
        }
        a, e := f(s)
        if e != nil {
            log.Fatal(e)
        }
        as = append(as, a)
    }
    return as
}

func mapIntSlice(ns []int, f func(int) IntPair) []IntPair {
    rs := make([]IntPair, len(ns))
    for i, n := range ns {
        rs[i] = f(n)
    }
    return rs
}

type IntPair struct {
    Left int
    Right int
}

func calculateWait(i, target int) int {
    base := i
    for i < target {
        i += base
    }
    return i - target
}

func minSlicePair(ps []IntPair) IntPair {
    min := ps[0]
    for _, i := range ps {
        if i.Right < min.Right {
            min = i
        }
    }
    return min
}

func main() {
    is := readInput("input")
    earliestDepartureTime, _ := strconv.Atoi(is[0])
    busses := mapStringSlice(strings.Split(is[1], ","), strconv.Atoi)
    var modTime = func(i int) IntPair {
        return IntPair{i, calculateWait(i, earliestDepartureTime)}
    }
    times := mapIntSlice(busses, modTime)
    solution := minSlicePair(times)
    log.Printf("%v: %d", solution, solution.Left * solution.Right)

    cs := make([]Constraint, 0)
    for i, b := range strings.Split(is[1], ",") {
        if b != "x" {
            busID, _ := strconv.Atoi(b)
            cs = append(cs, Constraint{busID, i})
        }
    }
    sort.Slice(cs, func(i, j int) bool {
        return cs[i].busID > cs[j].busID
    })

    s2 := solve(cs[0:1])
    log.Printf("solve 2 : %d", s2)
    log.Printf("solution to part 2: %d", solve(cs))
}

type Constraint struct {
    busID int
    index int
}

func solve(cs []Constraint) int {
    // we assume cs is sorted, that makes this faster
    index := 1
    current := cs[0].busID * index + cs[0].index // a candidate because it satisfies the first constraint
    for !satisfiesAll(current, cs) {
        index++
        current = cs[0].busID * index + cs[0].index
    }
    return current
}

func satisfiesAll(i int, cs []Constraint) bool {
    s := true
    for _, c := range cs {
        satisfied := satisfies(i, c)
        if !satisfied {
            return false
        }
    }
    return s
}

func satisfies(i int, c Constraint) bool {
    mod := int(math.Mod(float64(i +c.index), float64(c.busID)))
    return mod == 0
}