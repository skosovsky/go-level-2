package main

import "testing"

func Test_safeMutexAdd(t *testing.T) {
	tests := []struct {
		name string
	}{
		{""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			safeMutexAdd()
		})
	}
}

func Test_safeMutexDif(t *testing.T) {
	tests := []struct {
		name string
	}{
		{""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			safeMutexDif()
		})
	}
}
