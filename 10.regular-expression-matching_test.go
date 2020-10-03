package main

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		//	{"Example 1", args{"aa", "a"}, false},
		//	{"Example 2", args{"aa", "a*"}, true},
		//	{"Example 3", args{"ab", ".*"}, true},
		//	{"Example 4", args{"aab", "c*a*b"}, true},
		//	{"Example 5", args{"mississippi", "mis*is*p*.*"}, true},
		{"Wrong 1", args{"aaa", "ab*a*c*a"}, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
