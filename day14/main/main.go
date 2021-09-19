package main

import (
	"log"

	"github.com/jensdewaard/advent-of-code-2020/day14"
)

func main() {
	p := day14.ReadInput("day14/input")
	memory := day14.RunProgramOne(p)

	log.Printf("sum of memory %d\n", memory.Sum())

	memory = day14.RunProgramTwo(p)
	log.Printf("sum of memory2 %d\n", memory.Sum())
}
