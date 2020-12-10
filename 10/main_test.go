package main

import "testing"

func Test_arrangements(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{"single adapter", []int{3}, 1},
		{"no adapters", []int{}, 0},
		{"missing adapter", []int{0, 5}, 0},
		{"one way1", []int{0,1}, 1},
		{"one way2", []int{0,2}, 1},
		{"one way3", []int{0,3}, 1},
		{"one way6", []int{0,3,6}, 1},
		{"bridge adapter", []int{0, 2, 3}, 2},
		{"bridge adapter2", []int{0, 2, 3, 6}, 2},
		{"bridge adapter3", []int{0, 2, 3, 6, 9}, 2},
		{"bridge adapter3", []int{0, 2, 3, 6, 7}, 2},
		{"one skip possible", []int{0, 1, 2}, 2},
		{"two skips possible", []int{0,1, 2, 3}, 4},
		{"two skips possible twice", []int{0,1, 2, 3, 6, 7, 8, 9}, 16},		
		{"riks case aangevuld", []int{0, 1, 3, 4, 5, 6, 7}, 18},
		{"riks casus 2", []int{0, 1, 2, 3, 4}, 7},
		{"given test case 1", []int{0, 1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19}, 8},
		{"given test case 2", []int{0, 1, 2, 3, 4, 7, 8, 9, 10, 11, 14, 17, 18, 19, 20, 23, 24, 25, 28, 31, 32, 33, 34, 35, 38, 39, 42, 45, 46, 47, 48, 49}, 19208},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			memoization = make(map[int]int, 0)
			if got := arrangements(tt.args); got != tt.want {
				t.Errorf("arrangements() = %v, want %v", got, tt.want)
			}
		})
	}
}
