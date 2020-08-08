package main

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{"abcabcbb"}, 3},
		{"Example 2", args{"bbbbb"}, 1},
		{"Example 3", args{"pwwkew"}, 3},
		{"Wrong 1", args{"aa"}, 1},
		{"Wrong 2", args{"aabaab!bb"}, 3},
		{"Wrong 3", args{"au"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
