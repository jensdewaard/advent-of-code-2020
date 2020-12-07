package main

import (
	"fmt"
	"strconv"
    "strings"
    "io/ioutil"
    "log"
)

type edge struct {
    to *node
    weight int
}

type node struct {
    label string
    containedBy []edge
}

func sanitizeLabel(label string) string {
    s := strings.TrimSuffix(label, "bags")
    s = strings.TrimSuffix(s, "bag")
    s = strings.TrimSpace(s)
    if label == s {
        return s
    } else {
        return sanitizeLabel(s)
    }
}

func createNode(label string) *node {
    n := &node{}
    n.label = sanitizeLabel(label)
    n.containedBy = make([]edge, 0)
    nodes = append(nodes, n)
    return n
}

func (n *node) String() string {
    containedBy := ""
    for _, e := range n.containedBy {
        containedBy = containedBy + e.String()
    }
    return "{" + n.label + containedBy + "}"
}

func (e *edge) String() string {
    return "[" + e.to.label + "," + string(e.weight) + "]"
}

var nodes []*node

func (n *node) reachable() []*node {
    done := make([]*node, 0)
    todo := []*node{n}
    return reachableHelper(done, todo)
}

func isIn(n *node, list []*node) bool {
    for _, m := range list {
        if n.label == m.label {
            return true
        }
    }
    return false
}

func reachableHelper(done, todo []*node) []*node {
    if len(todo) == 0 {
        return done
    }
    curNode := todo[0]
    fmt.Printf("!!current node %v!!\n", curNode)
    
    for _, e := range curNode.containedBy {
        fmt.Printf("!! ede to %v\n", e.to)
        if !isIn(e.to, done) {
            done = append(done, e.to)
            todo = append(todo, e.to)
        }
    }
    return reachableHelper(done, todo[1:])
}

func retrieveNode(color string) *node {
    c := sanitizeLabel(color)
    for _, n := range nodes {
        if n.label == c {
            return n
        }
    }
    return nil
}

func parseLine(s string) *node {
    spl := strings.Split(s, "bags contain")
    curCol := sanitizeLabel(spl[0]) // bag color
    n := retrieveNode(curCol)
    if n == nil {
        // create node
        n = createNode(spl[0])
    }
    if strings.Contains(s, "no other") {
        return n
    }

    contains := strings.Split(spl[1], ",")
    for _, b := range contains {
        e := edge{}
        btxt := strings.TrimSpace(b)
        w, _ := strconv.Atoi(string(btxt[0]))
        e.weight = w //todo allow 2-digit numbers

        removeBag := strings.TrimSuffix(btxt[2:], "bag")
        removeBags := strings.TrimSuffix(removeBag, "bags")
        lbl := strings.TrimSpace(removeBags)
        tn := retrieveNode(lbl)
        if tn == nil {
            // create node
            tn = createNode(lbl)
        }
        e.to = n
        tn.containedBy = append(tn.containedBy, e)
    }
    return n
}

func main() {
    nodes = make([]*node, 0)
    rules := readInput("input")
    for _, l := range rules {
        parseLine(l)
    }
    fmt.Printf("%v\n", nodes)
    n := retrieveNode("shiny gold")
    fmt.Printf("%v\n", n.reachable())
}

func readInput(file string) []string {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	text := string(dat)
	return strings.Split(text, "\n")
}