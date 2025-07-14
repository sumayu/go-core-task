package handler

import (
	"reflect"
	"testing"
)

func TestAddElements(t *testing.T) {
	type args struct {
		slice []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "add to empty slice",
			args: args{
				slice: []int{},
				n:     5,
			},
			want: []int{5},
		},
		{
			name: "add to non-empty slice",
			args: args{
				slice: []int{1, 2, 3},
				n:     4,
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "add negative number",
			args: args{
				slice: []int{10, 20},
				n:     -5,
			},
			want: []int{10, 20, -5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddElements(tt.args.slice, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopySlice(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "copy empty slice",
			args: args{
				slice: []int{},
			},
			want: []int{},
		},
		{
			name: "copy single element",
			args: args{
				slice: []int{42},
			},
			want: []int{42},
		},
		{
			name: "copy multiple elements",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CopySlice(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CopySlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveElement(t *testing.T) {
	type args struct {
		slice []int
		index int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "remove first element",
			args: args{
				slice: []int{1, 2, 3},
				index: 0,
			},
			want: []int{2, 3},
		},
		{
			name: "remove middle element",
			args: args{
				slice: []int{10, 20, 30, 40},
				index: 2,
			},
			want: []int{10, 20, 40},
		},
		{
			name: "remove last element",
			args: args{
				slice: []int{5, 10, 15},
				index: 2,
			},
			want: []int{5, 10},
		},
		{
			name: "remove from single element slice",
			args: args{
				slice: []int{100},
				index: 0,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveElement(tt.args.slice, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceExample(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty slice",
			args: args{
				slice: []int{},
			},
			want: []int{},
		},
		{
			name: "all even numbers",
			args: args{
				slice: []int{2, 4, 6, 8},
			},
			want: []int{2, 4, 6, 8},
		},
		{
			name: "all odd numbers",
			args: args{
				slice: []int{1, 3, 5, 7},
			},
			want: []int{},
		},
		{
			name: "mixed even and odd numbers",
			args: args{
				slice: []int{1, 2, 3, 4, 5, 6},
			},
			want: []int{2, 4, 6},
		},
		{
			name: "single even number",
			args: args{
				slice: []int{10},
			},
			want: []int{10},
		},
		{
			name: "single odd number",
			args: args{
				slice: []int{11},
			},
			want: []int{},
		},
		{
			name: "with negative numbers",
			args: args{
				slice: []int{-2, -1, 0, 1, 2},
			},
			want: []int{-2, 0, 2},
		},
		{
			name: "large numbers",
			args: args{
				slice: []int{1000000, 1000001},
			},
			want: []int{1000000},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceExample(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceExample() = %v, want %v", got, tt.want)
			}
		})
	}
}
