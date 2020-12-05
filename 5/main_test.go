package main

import (
	"reflect"
	"testing"
)

func TestPower(t *testing.T) {
	type args struct {
		x int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{2, 4}, 16},
		{"2", args{2, 0}, 1},
		{"3", args{3, 2}, 9},
		{"4", args{2, 2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Power(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("Power() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassToInt(t *testing.T) {
	type args struct {
		bp BoardingPass
	}
	tests := []struct {
		name string
		args args
		want Seat
	}{
		{"1", args{"BFFFBBFRRR"}, Seat{70, 7}},
		{"2", args{"FFFBBBFRRR"}, Seat{14, 7}},
		{"3", args{"BBFFBBFRLL"}, Seat{102, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PassToInt(tt.args.bp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PassToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeat_getID(t *testing.T) {
	type fields struct {
		Row int
		Col int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"1", fields{70, 7}, 567},
		{"1", fields{14, 7}, 119},
		{"1", fields{102, 4}, 820},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Seat{
				Row: tt.fields.Row,
				Col: tt.fields.Col,
			}
			if got := s.getID(); got != tt.want {
				t.Errorf("Seat.getID() = %v, want %v", got, tt.want)
			}
		})
	}
}
