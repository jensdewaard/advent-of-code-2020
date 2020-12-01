package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type input []int

func sortInput(in []int) []int {
	sort.Ints(in)
	return in
}

var target = 2020

func selectPair(list []int) (int, int) {
Nexti:
	for i := 0; i < len(list) - 1; i++ {	
		for j := 1; j < len(list); j++ {
			if list[i] + list[j] == target {
				return list[i], list[j]
			} else if list[i] + list[j] > target {
				break Nexti
			}
		}
	}
	log.Fatal("no sum adds to target")
	return 0,0
}

func selectThree(list[] int) (int, int, int) {
	for i := 0; i < len(list) - 2; i++ {
		for j := 1; j < len(list) - 1; j++ {
			for k := 2; k < len(list) - 2; k++ {
				if list[i] + list[j] + list[k] == target {
					return list[i], list[j], list[k]
				}
			}
		}
	}
	log.Fatal("no sum adds to target")
	return 0,0,0
}

func mul(a, b int) int {
	return a * b
}

func mul3(a, b, c int) int {
	return a * b * c
}

func parseStrings(in []string) []int {
	out := []int{}
	for _, s := range(in) {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		out = append(out, i)
	}
	return out
}

func readInput(file string) []int {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	text := string(dat)
	lines := strings.Split(text, "\n")
	return parseStrings(lines)
}

func main() {
	inputData := readInput("input")
	fmt.Println(mul3(selectThree(sortInput(inputData))))
}