package main

import (
	"reflect"
	"testing"
)

func Test_turn(t *testing.T) {
	type args struct {
		p      Position
		degree int
	}
	tests := []struct {
		name string
		args args
		want Position
	}{
		{"turn left", args{Position{0, 0, 50}, -50}, Position{0, 0, 0}},
		{"no turn", args{Position{0, 0, 50}, 0}, Position{0, 0, 50}},
		{"no turn negative", args{Position{0, 0, 50}, -0}, Position{0, 0, 50}},
		{"turn over 0, left", args{Position{0, 0, 90}, -100}, Position{0, 0, 350}},
		{"turn over 0, right", args{Position{0, 0, 180}, 270}, Position{0, 0, 90}},
		{"turn circle left", args{Position{0, 0, 180}, -360}, Position{0, 0, 180}},
		{"turn circle right", args{Position{0, 0, 180}, 360}, Position{0, 0, 180}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := turn(tt.args.p, tt.args.degree); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("turn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManhattan(t *testing.T) {
	type args struct {
		p Position
		q Position
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"zero", args{Position{0, 0, 0}, Position{0, 0, 0}}, 0},
		{"(3,0) (0,3)", args{Position{3, 0, 0}, Position{0, 3, 0}}, 6},
		{"(0,0) (10,10)", args{Position{0, 0, 0}, Position{10, 10, 0}}, 20},
		{"(0,0) (17, -8)", args{Position{0, 0, 90}, Position{17, -8, 180}}, 25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Manhattan(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("Manhattan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoute(t *testing.T) {
	type args struct {
		as     []Action
		curPos Position
	}
	tests := []struct {
		name string
		args args
		want Position
	}{
		{"given test case", args{
			[]Action{
				Action{"F", 10},
				Action{"N", 3},
				Action{"F", 7},
				Action{"R", 90},
				Action{"F", 11},
			},
			Position{0, 0, 90},
		}, Position{17, -8, 180}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Route(tt.args.as, tt.args.curPos); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Route() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_turnWaypoint(t *testing.T) {
	type args struct {
		shipPos Position
		wayPos  Position
		degrees int
	}
	tests := []struct {
		name string
		args args
		want Position
	}{
		{"no turn", args{Position{0, 0, 0}, Position{1, 4, 0}, 0}, Position{1, 4, 0}},
		{"180 turn around origin", args{Position{0, 0, 0}, Position{1, 4, 0}, 180}, Position{-1, -4, 0}},
		{"180 turn", args{Position{3, 3, 0}, Position{3, 6, 0}, 180}, Position{3, 0, 0}},
		{"90 turn right", args{Position{3, 3, 0}, Position{3, 6, 0}, 90}, Position{6, 3, 0}},
		{"90 turn left", args{Position{0, 0, 0}, Position{1, 1, 0}, -90}, Position{-1, 1, 0}},
		{"90 turn right", args{Position{0, 0, 0}, Position{1, 1, 0}, 90}, Position{1, -1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := turnWaypoint(tt.args.shipPos, tt.args.wayPos, tt.args.degrees); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("turnWaypoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouteWaypoint(t *testing.T) {
	type args struct {
		as      []Action
		shipPos Position
		wayPos  Position
	}
	tests := []struct {
		name string
		args args
		want Position
	}{
		// TODO: Add test cases.
		{"given test case", args{
			[]Action{
				Action{"F", 10},
				Action{"N", 3},
				Action{"F", 7},
				Action{"R", 90},
				Action{"F", 11},
			},
			Position{0, 0, 90},
			Position{10, 1, 90},
		}, Position{214, -72, 90}},
		{"given test caseA", args{
			[]Action{
				Action{"F", 10},
			},
			Position{0, 0, 90},
			Position{10, 1, 90},
		}, Position{100, 10, 90}},
		{"given test caseB", args{
			[]Action{
				Action{"F", 10},
				Action{"N", 3},
			},
			Position{0, 0, 90},
			Position{10, 1, 90},
		}, Position{100, 10, 90}},
		{"given test caseC", args{
			[]Action{
				Action{"F", 10},
				Action{"N", 3},
				Action{"F", 7},
			},
			Position{0, 0, 90},
			Position{10, 1, 90},
		}, Position{170, 38, 90}},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RouteWaypoint(tt.args.as, tt.args.shipPos, tt.args.wayPos); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RouteWaypoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
