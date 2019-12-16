package main

import "testing"

func BenchmarkBinaryQueueConvert(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BinaryQueueConvert(400)
	}
}

func BenchmarkBinaryQueueDisplay(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BinaryQueueDisplay(400)
	}
}
