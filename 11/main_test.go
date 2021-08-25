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
		{"out of bounds x >", args{[]string{"...", "...", "..."}, 5, 2}, 'x'},
		{"out of bounds x <", args{[]string{"...", "...", "..."}, -1, 2}, 'x'},
		{"out of bounds y >", args{[]string{"...", "...", "..."}, 2, 9}, 'x'},
		{"out of bounds y <", args{[]string{"...", "...", "..."}, 1, -3}, 'x'},
		{"no lines", args{[]string{}, 1, 3}, 'x'},
		{"empty strings", args{[]string{"", "", ""}, 1, 3}, 'x'},
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

func Test_adjacentSeatsOccupiedTwo(t *testing.T) {
	type args struct {
		world []string
		x     int
		y     int
	}
	worldOne := []string{
		".......#.",
		"...#.....",
		".#.......",
		".........",
		"..#L....#",
		"....#....",
		".........",
		"#........",
		"...#.....",
	}
	worldTwo := []string{
		".............",
		".L.L.#.#.#.#.",
		".............",
	}
	worldThree := []string{
		".##.##.",
		"#.#.#.#",
		"##...##",
		"...L...",
		"##...##",
		"#.#.#.#",
		".##.##.",
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"example1", args{worldOne, 3, 4}, 8},
		{"example2a", args{worldTwo, 1, 1}, 0},
		{"example2b", args{worldTwo, 3, 1}, 1},
		{"example3", args{worldThree, 3, 3}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := adjacentSeatsOccupiedTwo(tt.args.world, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("adjacentSeatsOccupied() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_roundRuleOne(t *testing.T) {
	tests := []struct {
		name        string
		world       []string
		wantWorld   []string
		wantChanged bool
	}{
		{"emptySeatFilled", []string{"...", ".L.", "..."}, []string{"...", ".#.", "..."}, true},
		{"seatBecomesEmpty", []string{"#.#", ".#.", "###"}, []string{"#.#", ".L.", "###"}, true},
		{"seatRemainsEmpty", []string{".#.", "L.L", "..."}, []string{".#.", "L.L", "..."}, false},
		{
			"example1",
			[]string{"L.LL.LL.LL", "LLLLLLL.LL", "L.L.L..L..", "LLLL.LL.LL", "L.LL.LL.LL", "L.LLLLL.LL", "..L.L.....", "LLLLLLLLLL", "L.LLLLLL.L", "L.LLLLL.LL"},
			[]string{"#.##.##.##", "#######.##", "#.#.#..#..", "####.##.##", "#.##.##.##", "#.#####.##", "..#.#.....", "##########", "#.######.#", "#.#####.##"},
			true,
		},
		{
			"example2",
			[]string{"#.##.##.##", "#######.##", "#.#.#..#..", "####.##.##", "#.##.##.##", "#.#####.##", "..#.#.....", "##########", "#.######.#", "#.#####.##"},
			[]string{"#.LL.L#.##", "#LLLLLL.L#", "L.L.L..L..", "#LLL.LL.L#", "#.LL.LL.LL", "#.LLLL#.##", "..L.L.....", "#LLLLLLLL#", "#.LLLLLL.L", "#.#LLLL.##"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newWorld, changed := round(tt.world, ruleOne)
			if !stringSliceEqual(newWorld, tt.wantWorld) || changed != tt.wantChanged {
				t.Errorf("round() = (%v, %v), want (%v, %v)", newWorld, changed, tt.wantWorld, tt.wantChanged)
			}
		})
	}
}

func Test_roundRuleTwo(t *testing.T) {
	tests := []struct {
		name        string
		world       []string
		wantWorld   []string
		wantChanged bool
	}{
		{
			"example1",
			[]string{"L.LL.LL.LL", "LLLLLLL.LL", "L.L.L..L..", "LLLL.LL.LL", "L.LL.LL.LL", "L.LLLLL.LL", "..L.L.....", "LLLLLLLLLL", "L.LLLLLL.L", "L.LLLLL.LL"},
			[]string{"#.##.##.##", "#######.##", "#.#.#..#..", "####.##.##", "#.##.##.##", "#.#####.##", "..#.#.....", "##########", "#.######.#", "#.#####.##"},
			true,
		},
		{
			"example2",
			[]string{
                "#.##.##.##", 
                "#######.##", 
                "#.#.#..#..", 
                "####.##.##", 
                "#.##.##.##", 
                "#.#####.##", 
                "..#.#.....", 
                "##########", 
                "#.######.#", 
                "#.#####.##",
            },
			[]string{
                "#.LL.LL.L#", 
                "#LLLLLL.LL", 
                "L.L.L..L..", 
                "LLLL.LL.LL", 
                "L.LL.LL.LL", 
                "L.LLLLL.LL", 
                "..L.L.....", 
                "LLLLLLLLL#", 
                "#.LLLLLL.L", 
                "#.LLLLL.L#",
            },
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newWorld, changed := round(tt.world, ruleTwo)
			if !stringSliceEqual(newWorld, tt.wantWorld) || changed != tt.wantChanged {
				t.Errorf("round() = (%v, %v), want (%v, %v)", newWorld, changed, tt.wantWorld, tt.wantChanged)
			}
		})
	}
}

func Test_countSeats(t *testing.T) {
	tests := []struct {
		name  string
		world []string
		want  int
	}{
		{"5 seats", []string{".L..", "L##L", "##.#"}, 5},
		{"12 seats", []string{"####", "####", "####"}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countSeats(tt.world)
			if got != tt.want {
				t.Errorf("countSeat() = %d, want %d", got, tt.want)
			}
		})
	}
}

func stringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
