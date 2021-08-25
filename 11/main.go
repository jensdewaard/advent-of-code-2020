package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func readInput(file string) []string {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	text := string(dat)
	return strings.Split(text, "\n")
}

func main() {
	world := readInput("input")

	// part one
	changed := true
	for changed {
		world, changed = round(world, ruleOne)
	}
	log.Printf("1: %d seats occupied\n", countSeats(world))

	// part two
	world = readInput("input")

	changed = true
	for changed {
		world, changed = round(world, ruleTwo)
	}
	log.Printf("2: %d seats occupied\n", countSeats(world))

}

func countSeats(world []string) int {
	count := 0
	for _, l := range world {
		count += strings.Count(l, "#")
	}
	return count
}

const (
	OpenSeat     = 'L'
	OccupiedSeat = '#'
	Floor        = '.'
	OutOfBounds  = 'x'
)

type seatingRule func([]string, int, int) (rune, bool)

func ruleOne(world []string, w, h int) (rune, bool) {
	switch getSeatType(world, w, h) {
	case OpenSeat:
		if adjacentSeatsOccupied(world, w, h) == 0 {
			return OccupiedSeat, true
		} else {
			return OpenSeat, false
		}
	case OccupiedSeat:
		if adjacentSeatsOccupied(world, w, h) >= 4 {
			return OpenSeat, true
		} else {
			return OccupiedSeat, false
		}
	default:
		return Floor, false
	}
}

func ruleTwo(world []string, w, h int) (rune, bool) {
	switch getSeatType(world, w, h) {
	case OpenSeat:
		if adjacentSeatsOccupiedTwo(world, w, h) == 0 {
			return OccupiedSeat, true
		} else {
			return OpenSeat, false
		}
	case OccupiedSeat:
		if adjacentSeatsOccupiedTwo(world, w, h) >= 5 {
			return OpenSeat, true
		} else {
			return OccupiedSeat, false
		}
	default:
		return Floor, false
	}
}

func round(world []string, next seatingRule) (newWorld []string, changed bool) {
	changed = false
	newWorld = make([]string, len(world))
	for i := 0; i < len(world); i++ {
		newWorld[i] = strings.Repeat(".", len(world[0]))
	}
	for h := 0; h < len(world); h++ {
		row := ""
		for w := 0; w < len(world[0]); w++ {
			nextRune, change := next(world, w, h)
			row += string(nextRune)
			changed = changed || change
		}
		newWorld[h] = row
	}
	return
}

func getSeatType(world []string, x int, y int) rune {
	if y < 0 || y >= len(world) {
		return OutOfBounds
	}
	line := world[y]
	if x < 0 || x >= len(line) {
		return OutOfBounds
	}
	return rune(line[x])
}

func adjacentSeatsOccupied(world []string, x int, y int) int {
	// top row
	occupied := 0
	if getSeatType(world, x-1, y-1) == OccupiedSeat {
		occupied++
	}
	if getSeatType(world, x, y-1) == OccupiedSeat {
		occupied++
	}
	if getSeatType(world, x+1, y-1) == OccupiedSeat {
		occupied++
	}
	if getSeatType(world, x-1, y) == OccupiedSeat {
		occupied++
	}
	if getSeatType(world, x+1, y) == OccupiedSeat {
		occupied++
	}
	if getSeatType(world, x-1, y+1) == OccupiedSeat {
		occupied++
	}
	if getSeatType(world, x, y+1) == OccupiedSeat {
		occupied++
	}
	if getSeatType(world, x+1, y+1) == OccupiedSeat {
		occupied++
	}
	return occupied
}

func adjacentSeatsOccupiedDir(world []string, x, y int, d Dir) int {
	switch getSeatType(world, x, y) {
	case OpenSeat, OutOfBounds:
		return 0
	case OccupiedSeat:
		return 1
	default:
		switch d {
		case UPLEFT:
			return adjacentSeatsOccupiedDir(world, x-1, y-1, UPLEFT)
		case UP:
			return adjacentSeatsOccupiedDir(world, x, y-1, UP)
		case UPRIGHT:
			return adjacentSeatsOccupiedDir(world, x+1, y-1, UPRIGHT)
		case LEFT:
			return adjacentSeatsOccupiedDir(world, x-1, y, LEFT)
		case RIGHT:
			return adjacentSeatsOccupiedDir(world, x+1, y, RIGHT)
		case DOWNLEFT:
			return adjacentSeatsOccupiedDir(world, x-1, y+1, DOWNLEFT)
		case DOWN:
			return adjacentSeatsOccupiedDir(world, x, y+1, DOWN)
		case DOWNRIGHT:
			return adjacentSeatsOccupiedDir(world, x+1, y+1, DOWNRIGHT)
		}
		return 0
	}
}

func adjacentSeatsOccupiedTwo(world []string, x, y int) int {
	return adjacentSeatsOccupiedDir(world, x-1, y-1, UPLEFT) +
		adjacentSeatsOccupiedDir(world, x, y-1, UP) +
		adjacentSeatsOccupiedDir(world, x+1, y-1, UPRIGHT) +
		adjacentSeatsOccupiedDir(world, x-1, y, LEFT) +
		adjacentSeatsOccupiedDir(world, x+1, y, RIGHT) +
		adjacentSeatsOccupiedDir(world, x-1, y+1, DOWNLEFT) +
		adjacentSeatsOccupiedDir(world, x, y+1, DOWN) +
		adjacentSeatsOccupiedDir(world, x+1, y+1, DOWNRIGHT)
}

type Dir int

const (
	UPLEFT    Dir = 1
	UP        Dir = 2
	UPRIGHT   Dir = 3
	LEFT      Dir = 4
	RIGHT     Dir = 5
	DOWNLEFT  Dir = 6
	DOWN      Dir = 7
	DOWNRIGHT Dir = 8
)
