package main

import (
	"testing"
)

func Test_calculateWait(t *testing.T) {
	type args struct {
		i      int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"example given", args{59, 939}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateWait(tt.args.i, tt.args.target); got != tt.want {
				t.Errorf("calculateWait() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_satisfiesAll(t *testing.T) {
	type args struct {
		i  int
		cs []Constraint
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"given example", args{1068781, []Constraint{
			Constraint{7, 0},
			Constraint{13, 1},
			Constraint{59, 4},
			Constraint{31, 6},
			Constraint{19, 7},
		}}, true},
	}	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := satisfiesAll(tt.args.i, tt.args.cs); got != tt.want {
				t.Errorf("satisfiesAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
