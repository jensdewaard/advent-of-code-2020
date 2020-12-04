package main

import (
	"testing"
)

func Test_validYear(t *testing.T) {
	type args struct {
		y string
		l int
		u int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"1920", 1920, 2000}, true},
		{"2", args{"190", 1920, 2000}, false},
		{"3", args{"190yy", 1920, 2000}, false},
		{"4", args{"a2003y", 1920, 2005}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validYear(tt.args.y, tt.args.l, tt.args.u); got != tt.want {
				t.Errorf("validYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validHeight(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"150cm"}, true},
		{"2", args{"63in"}, true},
		{"3", args{"193cm"}, true},
		{"4", args{"200cm"}, false},
		{"5", args{"garbage"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validHeight(tt.args.v); got != tt.want {
				t.Errorf("validHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validHairColor(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"mixed hex", args{"#1234ab"}, true},
		{"all text", args{"#ffabab"}, true},
		{"all numbers", args{"#123456"}, true},
		{"no hashtag", args{"123456"}, false},
		{"wrong letters", args{"#afbrta"}, false},
		{"too long", args{"#0000000"}, false},
		{"too short", args{"#0000"}, false},
		{"garbage", args{"garbage"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validHairColor(tt.args.v); got != tt.want {
				t.Errorf("validHairColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validEyecolor(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"blue", args{"blu"}, true},
		{"hazel", args{"hzl"}, true},
		{"amber", args{"amb"}, true},
		{"brown", args{"brn"}, true},
		{"gray", args{"gry"}, true},
		{"green", args{"grn"}, true},
		{"other", args{"oth"}, true},
		{"red", args{"red"}, false},
		{"garbage", args{"garbage"}, false},
		{"blue2", args{"blue"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validEyecolor(tt.args.v); got != tt.want {
				t.Errorf("validEyecolor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validField(t *testing.T) {
	type args struct {
		f field
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"cid", args{ field {"cid", "12345"}}, true},
		{"ecl", args{ field {"ecl", "blu"}}, true},
		{"hcl", args{ field {"hcl", "#123456"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validField(tt.args.f); got != tt.want {
				t.Errorf("validField() = %v, want %v", got, tt.want)
			}
		})
	}
}
