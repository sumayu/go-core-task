package main

import (
	"testing"
)

func Test_generate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "simple sequence",
			nums: []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "empty sequence",
			nums: []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := generate(tt.nums...)
			var got []int

			for v := range ch {
				got = append(got, v)
			}

			if len(got) != len(tt.want) {
				t.Errorf("generate() length = %d, want %d", len(got), len(tt.want))
				return
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("generate() at index %d = %d, want %d", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func Test_redesign(t *testing.T) {
	inCh := make(chan int)
	outCh := make(chan uint8)
	resultCh := redesign(inCh, outCh)

	go func() {
		inCh <- 1
		inCh <- 2
		inCh <- 3
		close(inCh)
	}()

	expected := []uint8{1, 2, 3}
	var got []uint8

	for v := range resultCh {
		got = append(got, v)
	}

	if len(got) != len(expected) {
		t.Errorf("redesign() length = %d, want %d", len(got), len(expected))
		return
	}

	for i := range got {
		if got[i] != expected[i] {
			t.Errorf("redesign() at index %d = %d, want %d", i, got[i], expected[i])
		}
	}
}

func Test_redesignFloat(t *testing.T) {
	inCh := make(chan uint8)
	outCh := make(chan float64)
	resultCh := redesignFloat(inCh, outCh)

	go func() {
		inCh <- 1
		inCh <- 2
		inCh <- 3
		close(inCh)
	}()

	expected := []float64{1.0, 8.0, 27.0}
	var got []float64

	for v := range resultCh {
		got = append(got, v)
	}

	if len(got) != len(expected) {
		t.Errorf("redesignFloat() length = %d, want %d", len(got), len(expected))
		return
	}

	for i := range got {
		if got[i] != expected[i] {
			t.Errorf("redesignFloat() at index %d = %f, want %f", i, got[i], expected[i])
		}
	}
}
