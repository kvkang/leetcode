package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildListNode(i int) *ListNode {
	var l ListNode
	var atListNode = &l

	for i > 0 {
		atListNode.Val = i % 10
		i = i / 10
		if i > 0 {
			atListNode.Next = &ListNode{}
			atListNode = atListNode.Next
		}
	}
	return &l
}

func ListNodeToInt(l *ListNode) int {
	var deep int
	var i int
	for l != nil {
		if deep == 0 {
			i = l.Val
			deep = 10
		} else {
			i += l.Val * deep
			deep = deep * 10
		}
		l = l.Next
	}
	return i
}

func TestAddTwoNumbers(t *testing.T) {
	var args = []struct {
		l1     *ListNode
		l2     *ListNode
		expect *ListNode
	}{}

	for _, arg := range args {
		got := addTwoNumbers(arg.l1, arg.l2)

		if !reflect.DeepEqual(got, arg.expect) {
			t.Fatal(got, arg.expect)
		}
	}
}

// @lc code=end

func TestListNodeToInt(t *testing.T) {
	type args struct {
		l *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"ListNode to 0", args{&ListNode{}}, 0},
		{"ListNode to 1 level", args{&ListNode{Val: 1}}, 1},
		{"ListNode to 2 level", args{&ListNode{Val: 1, Next: &ListNode{Val: 3}}}, 31},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListNodeToInt(tt.args.l); got != tt.want {
				t.Errorf("ListNodeToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildListNode(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"ListNode to 0", args{0}, &ListNode{}},
		{"ListNode to 1 level", args{1}, &ListNode{Val: 1}},
		{"ListNode to 2 level", args{31}, &ListNode{Val: 1, Next: &ListNode{Val: 3}}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildListNode(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
	}

	for randTestN := 1; randTestN <= 20; randTestN++ {
		randL1, randL2 := rand.Intn(1000), rand.Intn(1000)
		sumRandL := randL1 + randL2
		tests = append(tests, struct {
			name string
			args args
			want *ListNode
		}{
			fmt.Sprintf(
				"rand test %d %d+%d-->%d", randTestN, randL1, randL2, sumRandL),
			args{buildListNode(randL1), buildListNode(randL2)}, buildListNode(sumRandL)})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v %d, want %v %d", got, ListNodeToInt(got), tt.want, ListNodeToInt(tt.want))
			}
		})
	}
}
