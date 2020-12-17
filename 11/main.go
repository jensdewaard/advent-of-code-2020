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
    
    changed := true
    for changed {
        world, changed = round(world)
    }
    
    for _, l := range world {
        log.Println(l)
    }
    log.Printf("%d seats occupied\n", countSeats(world))
}

func countSeats(world []string) int {
    count := 0
    for _, l := range world {
        count += strings.Count(l, "#")
    }
    return count
}

const (
    OpenSeat = 'L'
    OccupiedSeat = '#'
    Floor = '.'
)

func round(world []string) (newWorld []string, changed bool) {
    changed = false
    newWorld = make([]string, len(world))
    for i := 0; i < len(world); i++ {
        newWorld[i] = strings.Repeat(".", len(world[0]))
    }
    for h := 0; h < len(world); h++ {
        row := ""
        for w := 0; w < len(world[0]); w++ {
            
            switch getSeatType(world, w, h) {
            case 'L':
                if adjacentSeatsOccupied(world, w, h) == 0 {
                    row += "#"
                    changed = true
                } else {
                    row += "L"
                }
            case '#':
                if adjacentSeatsOccupied(world, w, h) >= 4 {
                    row += "L"
                    changed = true
                } else {
                    row += "#"
                }
            case '.':
                row += "."
            }
        }
        newWorld[h] = row
    }
    
    return
}

func getSeatType(world []string, x int, y int) rune {
    if y < 0 || y >= len(world) {
        return Floor
    }
    line := world[y]
    if x < 0 || x >= len(line) {
        return Floor
    }
    return rune(line[x])
}

func adjacentSeatsOccupied(world []string, x int, y int) int {
    // top row
    occupied := 0
    if getSeatType(world, x - 1, y - 1) == OccupiedSeat {
        occupied++
    }    
    if getSeatType(world, x, y - 1) == OccupiedSeat {
        occupied++
    }
    if getSeatType(world, x + 1, y - 1) == OccupiedSeat {
        occupied++
    }
    if getSeatType(world, x - 1, y) == OccupiedSeat {
        occupied++
    }
    if getSeatType(world, x + 1, y) == OccupiedSeat {
        occupied++
    }
    if getSeatType(world, x - 1, y + 1) == OccupiedSeat {
        occupied++
    }
    if getSeatType(world, x, y + 1) == OccupiedSeat {
        occupied++
    }
    if getSeatType(world, x + 1, y + 1) == OccupiedSeat {
        occupied++
    }
    return occupied
}