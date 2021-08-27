package main

import (
	"testing"
)

func TestBuffer(t *testing.T) {
	tests := map[string]struct {
		batch1       [][]byte
		batch2       [][]byte
		expectChange bool
	}{
		"no writes -> no change": {
			expectChange: false,
		},
		"writes then no writes -> change": {
			batch1: [][]byte{
				[]byte("foo"),
				[]byte("bar"),
			},
			expectChange: true,
		},
		"no writes then writes -> change": {
			batch2: [][]byte{
				[]byte("foo"),
				[]byte("bar"),
			},
			expectChange: true,
		},
		"same writes -> no change": {
			batch1: [][]byte{
				[]byte("foo"),
				[]byte("bar"),
			},
			batch2: [][]byte{
				[]byte("foo"),
				[]byte("bar"),
			},
			expectChange: false,
		},
		"same writes different sizes -> no change": {
			batch1: [][]byte{
				[]byte("foo"),
				[]byte("bar"),
			},
			batch2: [][]byte{
				[]byte("f"),
				[]byte("o"),
				[]byte("o"),
				[]byte("bar"),
			},
			expectChange: false,
		},
		"short write then long write -> change": {
			batch1: [][]byte{
				[]byte("foo"),
			},
			batch2: [][]byte{
				[]byte("foo"),
				[]byte("foo"),
			},
			expectChange: true,
		},
		"long write then short write -> change": {
			batch1: [][]byte{
				[]byte("foo"),
				[]byte("foo"),
			},
			batch2: [][]byte{
				[]byte("foo"),
			},
			expectChange: true,
		},
		"different writes -> change": {
			batch1: [][]byte{
				[]byte("foo"),
			},
			batch2: [][]byte{
				[]byte("bar"),
			},
			expectChange: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var wrote int
			buf := &ChangeDetectingBuffer{}
			for _, bytes := range tt.batch1 {
				l, err := buf.Write(bytes)
				if err != nil {
					t.Errorf("expect no error, got %v", err)
				}
				if l != len(bytes) {
					t.Errorf("expect %v bytes written, got %v", len(bytes), l)
				}
				wrote += l
			}

			buf.Done()
			if wrote != 0 && !buf.Changed {
				t.Errorf("expect buf Changed, but it did not")
			}

			buf.Rewind()

			for _, bytes := range tt.batch2 {
				l, err := buf.Write(bytes)
				if err != nil {
					t.Errorf("expect no error, got %v", err)
				}
				if l != len(bytes) {
					t.Errorf("expect %v bytes written, got %v", len(bytes), l)
				}
				wrote += l
			}

			buf.Done()
			if tt.expectChange != buf.Changed {
				t.Errorf("expect changed to be %v, got %v", tt.expectChange, buf.Changed)
			}
		})
	}
}
