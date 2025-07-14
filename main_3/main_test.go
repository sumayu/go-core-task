package main

import (
	"reflect"
	"testing"
)

func TestCopy(t *testing.T) {
	type args struct {
		StringIntMap map[string]int
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "test_1",
			args: args{
				StringIntMap: map[string]int{"apple": 1, "banana": 2, "cherry": 3},
			},
			want: map[string]int{"apple": 1, "banana": 2, "cherry": 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Copy(tt.args.StringIntMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExists(t *testing.T) {
	type args struct {
		StringIntMap map[string]int
		key          string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test_1",
			args: args{
				StringIntMap: map[string]int{"apple": 1, "banana": 2, "cherry": 3},
				key:          "apple",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exists(tt.args.StringIntMap, tt.args.key); got != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_add(t *testing.T) {
	type args struct {
		StringIntMap map[string]int
		key          string
		value        int
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "test_1",
			args: args{
				StringIntMap: map[string]int{},
				key:          "apple",
				value:        5,
			},
			want: map[string]int{"apple": 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.StringIntMap, tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_remove(t *testing.T) {
	type args struct {
		StringIntMap map[string]int
		key          string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "test_1",
			args: args{
				StringIntMap: map[string]int{"apple": 1, "banana": 2, "cherry": 3},
				key:          "apple",
			},
			want: map[string]int{"banana": 2, "cherry": 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := remove(tt.args.StringIntMap, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("remove() = %v, want %v", got, tt.want)
			}
		})
	}
}
