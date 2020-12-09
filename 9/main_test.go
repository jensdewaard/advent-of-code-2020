package main

import "testing"

func Test_isSum(t *testing.T) {
	type args struct {
		is []int
		t  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"6yes", args{[]int{1, 2, 3, 4 ,5}, 6}, true},
		{"10no", args{[]int{1, 2, 3, 4 ,5}, 10}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSum(tt.args.is, tt.args.t); got != tt.want {
				t.Errorf("isSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
