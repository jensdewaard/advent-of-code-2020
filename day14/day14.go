package day14

import (
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func ReadInput(file string) Program {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	text := string(dat)
	ss := strings.Split(text, "\n")
	cs := make([]Command, 0)
	for _, s := range ss {
		cs = append(cs, parseCommand(s))
	}
	return cs
}

func parseCommand(s string) Command {
	if strings.HasPrefix(s, "mask") {
		m := strings.TrimPrefix(s, "mask = ")
		ma := MaskAssignment{Mask: Bitmask(m)}
		return ma
	} else if strings.HasPrefix(s, "mem") {
		ss := strings.Split(s, " = ")
		newValue, err := strconv.Atoi(ss[1])
		if err != nil {
			panic(err)
		}
		address, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(ss[0], "mem["), "]"))
		if err != nil {
			panic(err)
		}
		return MemoryInstruction{Address: address, Value: newValue}
	} else {
		panic("invalid input format")
	}
}

func Reverse(s Bitstring) Bitstring {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return Bitstring(string(runes))
}

func floatToBitstring(n float64) Bitstring {
	s := ""
	i := float64(35)
	for i >= 0 {
		if n >= math.Exp2(i) {
			s += "1"
			n -= math.Exp2(i)
		} else {
			s += "0"
		}
		i--
	}
	return Bitstring(s)
}

func BitstringToFloat64(bs Bitstring) float64 {
	br := Reverse(bs)
	result := float64(0)
	for i, b := range br {
		bit, _ := strconv.Atoi(string(b))
		result += float64(bit) * math.Exp2(float64(i))
	}
	return result
}

func maskString(v Bitstring, m Bitmask) Bitstring {
	result := ""
	for i, b := range m {
		if b == 'X' {
			result += string(v[i])
		} else {
			result += string(b)
		}
	}
	return Bitstring(result)
}

func maskAddress(v Bitstring, m Bitmask) []Bitstring {
	result := make([]Bitstring, 1)
	result[0] = Bitstring("")

	for i, b := range m {
		switch b {
		case 'X': // floating
			for _, r := range result {
				newRzero := r + "0"
				newRone := r + "1"
				result = append(result, newRzero)
				result = append(result, newRone)
			}
		case '0': //unchanged from v
			for _, r := range result {
				newR := r + Bitstring(v[i])
				r = newR
			}
		case '1': // output 1
			for _, r := range result {
				newR := r + "1"
				r = newR
			}
		}
	}
	return result
}

func RunProgramOne(p Program) Memory {
	var mask Bitmask = ""
	memory := make(Memory)
	for _, c := range p {
		switch c := c.(type) {
		case MaskAssignment:
			mask = c.Mask
		case MemoryInstruction:
			maskedValue := maskString(floatToBitstring(float64(c.Value)), mask)
			add := int(BitstringToFloat64(maskedValue))
			memory[c.Address] = float64(add)
		}
	}
	return memory
}
