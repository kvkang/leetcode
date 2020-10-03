package main

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name        string
		args        args
		want        int
		Explanation string
	}{
		{"Example 1", args{"III"}, 3, ""},
		{"Example 2", args{"IV"}, 4, ""},
		{"Example 3", args{"IX"}, 9, ""},
		{"Example 4", args{"LVIII"}, 58, "L = 50, V = 5 III = 3"},
		{"Example 5", args{"MCMXCIV"}, 1994, "M = 1000, CM = 900, XC = 90, IV = 4"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
