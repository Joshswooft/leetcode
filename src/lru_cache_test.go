package main

import (
	"reflect"
	"testing"
)

func TestNewLRUCache(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name   string
		args   args
		want   *lruCache
		expErr error
	}{
		{
			name: "invalid capacity",
			args: args{
				capacity: -2,
			},
			want:   nil,
			expErr: errInvalidCapacity,
		},
		{
			name: "makes cache with 4 capacity",
			args: args{
				capacity: 4,
			},
			want: &lruCache{
				capacity: 4,
			},
			expErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLRUCache(tt.args.capacity)
			if err != tt.expErr {
				t.Errorf("NewLRUCache() error = %v, wantErr %v", err, tt.expErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLRUCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
