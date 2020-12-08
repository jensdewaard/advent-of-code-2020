package main

import (
	"reflect"
	"testing"
)

func Test_parseInput(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want instruction
	}{
		{"noop", args{"nop +0"}, instruction{"nop", 0, false}},
		{"acc +5", args{"acc +5"}, instruction{"acc", 5, false}},
		{"jmp -4", args{"jmp -4"}, instruction{"jmp", -4, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseInput(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
