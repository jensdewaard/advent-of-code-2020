package main

import(
    "io/ioutil"
    "log"
    "strings"
    "strconv"
    "math"
)

type Action struct {
    dir string
    val int
}

func readInput(file string) []string {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	text := string(dat)
	return strings.Split(text, "\n")
}

type fn func(string) (Action, error)



func mapStringSlice(ss []string, f fn) []Action {
    as := make([]Action, 0)
    for _, s := range ss {
        a, e := f(s)
        if e != nil {
            log.Fatal(e)
        }
        as = append(as, a)
    }
    return as
}

func AtoAction(s string) (Action, error) {
    a := Action{}
    i, e := strconv.Atoi(s[1:])
    if e != nil {
        return a, e
    }
    a.dir = string(s[0])
    a.val = i
    return a, nil
}

type Position struct {
    X int
    Y int
    Facing int
}

func Manhattan(p, q Position) int {
    return int(math.Abs(float64(p.X - q.X)) + math.Abs(float64(p.Y - q.Y)))
}

func turn(p Position, degree int) Position {
    q := Position{}
    q.X = p.X
    q.Y = p.Y
    q.Facing = int(math.Mod(math.Mod(float64(p.Facing + degree), 360), -360))
    for q.Facing < 0 {
        q.Facing = 360 + q.Facing
    }
    return q
}

func Route(as []Action, curPos Position) Position {
    if len(as) == 0 {
        return curPos
    }
    a := as[0]
    newPos := Position{curPos.X, curPos.Y, curPos.Facing}
    switch a.dir {
    case "L": 
        d := -1 * a.val
        newPos = turn(curPos, d)
        
    case "R":
        newPos = turn(curPos, a.val)
        
    case "N":
        newPos.Y += a.val
        
    case "E":
        newPos.X += a.val
        
    case "W":
        newPos.X -= a.val
        
    case "S":
        newPos.Y -= a.val
    case "F":
        switch curPos.Facing {
        case 90:
            newPos.X += a.val
        case 180:
            newPos.Y -= a.val
        case 270:
            newPos.X -= a.val
        case 0:
            newPos.Y += a.val
        }
    }
    return Route(as[1:], newPos)
}

func main() {
    as := mapStringSlice(readInput("input"), AtoAction)
    start := Position{0,0,90}
    waypoint := Position{10, 1, 0}
    end := RouteWaypoint(as, start, waypoint)
    log.Printf("Manhattan %d", Manhattan(start, end))
}

func RouteWaypoint(as []Action, shipPos, wayPos Position) Position {
    if len(as) == 0 {
        return shipPos
    }
    a := as[0]
    newPosShip := Position{shipPos.X, shipPos.Y, shipPos.Facing}
    newPosWay := Position{wayPos.X, wayPos.Y, wayPos.Facing}
    switch a.dir {
    case "N":
        newPosWay.Y += a.val
        
    case "E":
        newPosWay.X += a.val
        
    case "W":
        newPosWay.X -= a.val
        
    case "S":
        newPosWay.Y -= a.val
    case "L": 
        d := -1 * a.val
        newPosWay = turnWaypoint(shipPos, wayPos, d)
        
    case "R":
        newPosWay = turnWaypoint(shipPos, wayPos, a.val)
    case "F":
        diffX := wayPos.X - shipPos.X
        diffY := wayPos.Y - shipPos.Y
        toMoveX := diffX * a.val
        toMoveY := diffY * a.val
        newPosShip.X += toMoveX
        newPosShip.Y += toMoveY
        newPosWay.X += toMoveX
        newPosWay.Y += toMoveY
        return RouteWaypoint(as[1:], newPosShip, newPosWay)
    }
    return RouteWaypoint(as[1:], newPosShip, newPosWay)
}

// 7189 is too low

func turnWaypoint(shipPos, wayPos Position, degrees int) Position {
    if degrees < 0 {
        degrees = -1 * degrees
        degrees = int(math.Mod(float64(degrees), 360))
        degrees = -1 * degrees
    } else {
        degrees = int(math.Mod(float64(degrees), 360))
    }
    if degrees == 0 {
        return wayPos
    }
    // translate waypoint pos wrt origin
    translatedWaypoint := Position{}
    translatedWaypoint.X = wayPos.X - shipPos.X
    translatedWaypoint.Y = wayPos.Y - shipPos.Y

    result := Position{}
    // the translated position can be rotated around origin by flipping signs
    // and switching x / y
    switch degrees {
    case 90, -270:
        result.X = translatedWaypoint.Y
        result.Y = -1 * translatedWaypoint.X
    case 180, -180:
        result.X = -1 * translatedWaypoint.X
        result.Y = -1 * translatedWaypoint.Y
    case 270, -90:
        result.X = -1 * translatedWaypoint.Y
        result.Y = translatedWaypoint.X
    }
    // translate the result back 
    result.X += shipPos.X
    result.Y += shipPos.Y

    return result

}