package main

import (
	"testing"
)

func Test_getSeatType(t *testing.T) {
	type args struct {
		world []string
		x     int
		y     int
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{"out of bounds x >", args{[]string{"...", "...", "..."}, 5, 2}, '.'},
		{"out of bounds x <", args{[]string{"...", "...", "..."}, -1, 2}, '.'},
		{"out of bounds y >", args{[]string{"...", "...", "..."}, 2, 9}, '.'},
		{"out of bounds y <", args{[]string{"...", "...", "..."}, 1, -3}, '.'},
		{"no lines", args{[]string{}, 1, 3}, '.'},
		{"empty strings", args{[]string{"", "", ""}, 1, 3}, '.'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSeatType(tt.args.world, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("getSeatType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_adjacentSeatsOccupied(t *testing.T) {
	type args struct {
		world []string
		x     int
		y     int
	}
	world := []string{
		".....",
		".L##L",
		".###L",
		".#.#.",
		".....",
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1,1", args{world, 1, 1}, 3},
		{"2,2", args{world, 2, 2}, 6},
		{"3,3", args{world, 3, 3}, 2},
		{"2,3", args{world, 3, 2}, 4},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := adjacentSeatsOccupied(tt.args.world, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("adjacentSeatsOccupied() = %v, want %v", got, tt.want)
			}
		})
	}
}
