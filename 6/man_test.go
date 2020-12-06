package main

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	type args struct {
		m1 map[byte]bool
		m2 map[byte]bool
	}
	tests := []struct {
		name string
		args args
		want map[byte]bool
	}{
		{
			"1", 
			args{
				map[byte]bool{
					'a': true,
					'b': true,
				},
				map[byte]bool{
					'b': true,
					'c': true,
				},
			},
			map[byte]bool{	
				'b': true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersect(tt.args.m1, tt.args.m2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
