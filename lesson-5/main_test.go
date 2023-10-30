package main

import "testing"

func BenchmarkMoreReadMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MoreReadMutex()
	}
}

func BenchmarkMoreReadRWMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MoreReadRWMutex()
	}
}
