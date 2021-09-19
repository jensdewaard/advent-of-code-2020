package main

import (
	"testing"
)

func Test_runProgramOne(t *testing.T) {
	tests := []struct {
		name  string
		input Program
		want  int
	}{
		{
			"example",
			Program{
				MaskAssignment{Bitmask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")},
				MemoryInstruction{8, 11},
				MemoryInstruction{7, 101},
				MemoryInstruction{8, 0},
			},
			165,
		},
		{
			"example_part1",
			Program{
				MaskAssignment{Bitmask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")},
				MemoryInstruction{8, 11},
			},
			73,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			memory := runProgramOne(tt.input)
			var result = float64(0)
			for _, v := range memory {
				result += v
			}

			got := int(result)
			if got != tt.want {
				t.Errorf("runProgram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_floatToBitstring(t *testing.T) {
	type args struct {
		n float64
	}
	tests := []struct {
		name string
		args args
		want Bitstring
	}{
		{"11", args{11}, "000000000000000000000000000000001011"},
		{"101", args{101}, "000000000000000000000000000001100101"},
		{"0", args{0}, "000000000000000000000000000000000000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := floatToBitstring(tt.args.n); got != tt.want {
				t.Errorf("iToBitstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	type args struct {
		s Bitstring
	}
	tests := []struct {
		name string
		args args
		want Bitstring
	}{
		{"test", args{Bitstring("test")}, "tset"},
		{"trigonometry", args{Bitstring("trigonometry")}, "yrtemonogirt"},
		{"foobar", args{Bitstring("foobar")}, "raboof"},
		{"1234", args{Bitstring("1234")}, "4321"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.s); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBitstringToFloat64(t *testing.T) {
	type args struct {
		bs Bitstring
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"101", args{"101"}, 5},
		{"1001", args{"1001"}, 9},
		{"111", args{"111"}, 7},
		{"0000", args{"0000"}, 0},
		{"0", args{"0"}, 0},
		{"00000", args{"00000"}, 0},
		{"000000000000000000000000000000001011", args{"000000000000000000000000000000001011"}, 11},
		{"000000000000000000000000000001001001", args{"000000000000000000000000000001001001"}, 73},
		{"000000000000000000000000000001100101", args{"000000000000000000000000000001100101"}, 101},
		{"000000000000000000000000000001000000", args{"000000000000000000000000000001000000"}, 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BitstringToFloat64(tt.args.bs); got != tt.want {
				t.Errorf("BitstringToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mask(t *testing.T) {
	type args struct {
		v Bitstring
		m Bitmask
	}
	tests := []struct {
		name string
		args args
		want Bitstring
	}{
		{"example 1", args{"000000000000000000000000000000001011", "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"}, "000000000000000000000000000001001001"},
		{"example 2", args{"000000000000000000000000000001100101", "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"}, "000000000000000000000000000001100101"},
		{"example 3", args{"000000000000000000000000000000000000", "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"}, "000000000000000000000000000001000000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maskString(tt.args.v, tt.args.m); got != tt.want {
				t.Errorf("mask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memorySum(t *testing.T) {
	tests := []struct {
		name string
		mem  Memory
		want int
	}{
		{
			"single value",
			map[int]float64{5: float64(20)},
			20,
		},
		{
			"two values",
			map[int]float64{1: float64(9), 9: float64(90)},
			99,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mem.Sum(); got != tt.want {
				t.Errorf("memory.Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maskAddress(t *testing.T) {
	tests := []struct {
		name    string
		address Bitstring
		mask    Bitmask
		want    []Bitstring
	}{
		{
			"example",
			Bitstring("000000000000000000000000000000101010"),
			Bitmask("000000000000000000000000000000X1001X"),
			[]Bitstring{
				"000000000000000000000000000000011010",
				"000000000000000000000000000000011011",
				"000000000000000000000000000000111010",
				"000000000000000000000000000000111011",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maskAddress(tt.address, tt.mask)
			if len(got) != len(tt.want) {
				t.Errorf("maskAddress() gives %d addresses, want %d", len(got), len(tt.want))
                return
			}
			for i, s := range got {
				if s != tt.want[i] {
					t.Errorf("maskAddress(%d) = %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}
