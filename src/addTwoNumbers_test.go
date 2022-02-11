package main

import (
	"reflect"
	"testing"
)

func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &ListNode{Val: nums[0], Next: nil}
	}
	remainingNums := nums[len(nums)-(len(nums)-1):]
	return &ListNode{Val: nums[0], Next: createLinkedList(remainingNums)}
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "list 1 and list 2 are empty",
			args: args{
				l1: []int{},
				l2: []int{},
			},
			want: []int{0},
		},
		{
			name: "example 1",
			args: args{
				l1: []int{2, 4, 3},
				l2: []int{5, 6, 4},
			},
			want: []int{7, 0, 8},
		},
		{
			name: "example 2",
			args: args{
				l1: []int{0},
				l2: []int{0},
			},
			want: []int{0},
		},
		{
			name: "example 3",
			args: args{
				l1: []int{9, 9, 9, 9, 9, 9, 9},
				l2: []int{9, 9, 9, 9},
			},
			want: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := createLinkedList(tt.args.l1)
			l2 := createLinkedList(tt.args.l2)

			got := addTwoNumbers(l1, l2)

			nums := got.ToList()

			if !reflect.DeepEqual(nums, tt.want) {
				t.Errorf("numbers not matching: got: %v, want: %v", nums, tt.want)
			}

		})
	}
}
