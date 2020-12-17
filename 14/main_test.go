package main

import (
	"testing"
)

func Test_iToBitstring(t *testing.T) {
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
			if got := iToBitstring(tt.args.n); got != tt.want {
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
		{"hidde", args{Bitstring("hidde")}, "eddih"},
		{"wiggert", args{Bitstring("wiggert")}, "treggiw"},
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
			if got := mask(tt.args.v, tt.args.m); got != tt.want {
				t.Errorf("mask() = %v, want %v", got, tt.want)
			}
		})
	}
}
