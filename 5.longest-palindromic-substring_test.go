package main

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Example 1", args{"babad"}, []string{"aba", "bab"}},
		{"Example 2", args{"cbbd"}, []string{"bb"}},
		{"Wrong 1", args{""}, []string{""}},
		{"Wrong 2", args{"a"}, []string{"a"}},
		{"Wrong 3", args{"abcda"}, []string{"a"}},
		{"Wrong 4", args{"ccc"}, []string{"ccc"}},
		// TODO: Add test cases.
	}

	inWants := func(got string, wants []string) bool {
		for _, want := range wants {
			if got == want {
				return true
			}
		}
		return false
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); !inWants(got, tt.want) {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
