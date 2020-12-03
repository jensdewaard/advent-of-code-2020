package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type vector struct {
    X,Y int
}

func (v vector) add(v2 vector) vector {
    return vector{v.X + v2.X, v.Y + v2.Y}
}

func at(location vector, terrain []string) rune {
    width := len(terrain[0])
    return rune(terrain[location.Y][location.X % width])
}

var slopes = []vector{
    vector{1, 1},
    vector{3, 1},
    vector{5, 1},
    vector{7, 1},
    vector{1, 2},
}

func main() {
    answers := make([]int, 5)
    for i, s := range slopes {
        count := 0
        currentLocation := vector{0,0}
        terrain := readInput("input")
        for currentLocation.Y < len(terrain) {
            if at(currentLocation, terrain) == '#' {
                count++
            }
            currentLocation = currentLocation.add(s)
        }
        fmt.Printf("%d: %d trees hit\n", i, count)
        answers[i] = count
    }
    fmt.Printf("Total multiplied %d\n", answers[0] * answers[1] * answers[2] * answers[3] * answers[4])
}

func readInput(file string) []string {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	text := string(dat)
	return strings.Split(text, "\n")
}
