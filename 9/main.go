package main

import(
    "io/ioutil"
    "log"
    "sort"
    "strings"
    "strconv"
)

var preambleSize = 25

var preamble []int
var preambleMin = 0
var preambleMax = 0

func preambleAdd(i int) int {
    l := preambleMax - preambleMin
    if l < preambleSize {
        // preamble is not yet filled
        preamble = append(preamble, i)
        preambleMax++
    } else {
        if !isSum(preamble[preambleMin:preambleMax], i) {
            log.Printf("%v", i)
            return i
        }
        preamble = append(preamble, i)
        preambleMin++
        preambleMax++
    }
    return 0
}

func isSum(is []int, t int) bool {
    cs := make([]int, preambleSize)
    copy(cs, is)
    sort.Ints(cs)
    for i := 0; i < preambleSize - 1; i++ {
        for j := i + 1; j < preambleSize; j++ {
            if cs[i] + cs[j] == t {
                return true
            }
        }
    }
    return false
}

func rangeSum(is []int, t int) bool {
    sum := 0
    for _, i := range is {
        sum += i
    }
    return sum == t
}

func minSlice(is []int) int {
    min := is[0]
    for _, i := range(is) {
        if i < min {
            min = i
        }
    }
    return min
}

func maxSlice(is []int) int {
    max := is[0]
    for _, i := range(is) {
        if i > max {
            max = i
        }
    }
    return max
}

func main() {
    ls := readInput("input")
    is := make([]int, 0)
    for _, l := range ls {
        i, err := strconv.Atoi(l)
        if err != nil {
            log.Fatal(err)
        }
        is = append(is, i)
    }
    //part 1
    magicNumber := 0
    for _, i := range is {
        magicNumber = preambleAdd(i)
        if magicNumber != 0 {
            break
        }
    }

    //part 2
    for i := 0; i < len(is) - 1; i++ {
        for j := i + 1; j < len(is); j++ {
            if rangeSum(is[i:j], magicNumber) {
                mi := minSlice(is[i:j])
                mx := maxSlice(is[i:j])
                log.Printf("%d + %d = %d", mi, mx, mi + mx)
            }
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