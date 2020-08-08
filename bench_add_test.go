package main

import "testing"

func AddOne(a int) int {
	var i, j int

	if a%2 == 0 {
		i = a
	}

	if a%3 == 0 {
		j = a
	}

	return a%4 + i + j
}

func AddMutil(a int) int {
	carry := a % 4
	if a%2 == 0 {
		carry += a
	}

	if a%3 == 0 {
		carry += a
	}

	return carry
}

func BenchmarkAddOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddOne(i)
	}
}

func BenchmarkAddMutil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddMutil(i)
	}
}
