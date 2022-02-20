package main

import (
	"fmt"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "finds full length palindrome",
			args: "rotator",
			want: "rotator",
		},
		{
			name: "example 1",
			args: "babad",
			want: "bab", // "aba" is also valid
		},
		{
			name: "example 2",
			args: "cbbd",
			want: "bb",
		},
		{
			name: "single letter",
			args: "a",
			want: "a",
		},
		{
			name: "not a palindrome",
			args: "ac",
			want: "c", // a is also valid
		},
		{
			name: "aacabdkacaa",
			args: "aacabdkacaa",
			want: "aca",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("test for: ", tt.args)
			if got := longestPalindrome(tt.args); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "reverse string",
			args: args{
				s: "one",
			},
			want: "eno",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.s); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
