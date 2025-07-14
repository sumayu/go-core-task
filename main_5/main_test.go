package main

import (
	"reflect"
	"testing"
)

func Test_intersection(t *testing.T) {
	type args struct {
		slice1 []int
		slice2 []int
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 []int
	}{
		{
			name: "test1",
			args: args{
				slice1: []int{65, 3, 58, 678, 64},
				slice2: []int{2, 43},
			},
			want:  false,
			want1: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := intersection(tt.args.slice1, tt.args.slice2)
			if got != tt.want {
				t.Errorf("intersection() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("intersection() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
