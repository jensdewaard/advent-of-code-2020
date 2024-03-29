package day14

type Program []Command
type Command interface {
	IsCommand()
}

type MemoryInstruction struct {
	Address int
	Value   int
}

func (MemoryInstruction) IsCommand() {}

type MaskAssignment struct {
	Mask Bitmask
}

func (MaskAssignment) IsCommand() {}

type Bitmask string

type Bitstring string

func (b Bitstring) append(c rune) Bitstring {
	s := string(b)
	return Bitstring(s + string(c))
}

type Memory map[int]float64

func (m Memory) Sum() int {
	result := float64(0)
	for _, v := range m {
		result += v
	}
	return int(result)
}
