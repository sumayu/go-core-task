package main

import "testing"

func TestMyWaitGroup_Add(t *testing.T) {
	type fields struct {
		counter int
		done    chan struct{}
	}
	type args struct {
		num int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Add",
			fields: fields{counter: 0, done: make(chan struct{})},
			args:   args{num: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mwg := &MyWaitGroup{
				counter: tt.fields.counter,
				done:    tt.fields.done,
			}
			mwg.Add(tt.args.num)
		})
	}
}

func TestMyWaitGroup_Done(t *testing.T) {
	type fields struct {
		counter int
		done    chan struct{}
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "Done",
			fields: fields{counter: 0, done: make(chan struct{})},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mwg := &MyWaitGroup{
				counter: tt.fields.counter,
				done:    tt.fields.done,
			}
			mwg.Done()
		})
	}
}
