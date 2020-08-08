package main

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{123}, 321},
		{"Example 2", args{-123}, -321},
		{"Example 2", args{120}, 21},
		{"Wrong 1", args{1534236469}, 0},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MeConfuse1(t *testing.T) {
	x := -12300
	for x != 0 {
		carry := x % 10
		t.Log(x, carry)
		x = x / 10
	}
}
