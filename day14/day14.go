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

func parseSingleMaskBit(v Bitstring, c, o rune) []Bitstring {
	switch c {
	case 'X': // floating
		return []Bitstring{
			v.append('0'),
			v.append('1'),
		}
	case '0':
		return []Bitstring{v.append(o)}
	case '1':
		return []Bitstring{v.append('1')}
	}
	return []Bitstring{}
}

func flatten(bss [][]Bitstring) []Bitstring {
	res := []Bitstring{}
	for _, bs := range bss {
		res = append(res, bs...)
	}
	return res
}

func maskAddress(v Bitstring, m Bitmask) []Bitstring {
	result := make([]Bitstring, 1)
	result[0] = Bitstring("")

	for i, b := range m {
		news := make([][]Bitstring, 0)
		for _, s := range result {
			news = append(news, parseSingleMaskBit(s, b, rune(v[i])))
		}
		result = flatten(news)
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

func RunProgramTwo(p Program) Memory {
	var mask Bitmask = ""
	memory := make(Memory)
	for _, c := range p {
		switch c := c.(type) {
		case MaskAssignment:
			mask = c.Mask
		case MemoryInstruction:
			maskedAddress := maskAddress(
				floatToBitstring(float64(c.Address)),
				mask,
			)
			for _, a := range maskedAddress {
				address := BitstringToFloat64(a)
				memory[int(address)] = float64(c.Value)
			}
		}
	}
	return memory
}
