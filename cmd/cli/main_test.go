package main

import (
	"sync"
	"testing"
)

func BenchmarkSliceDiff(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sliceDiff(flights)
	}
}

func BenchmarkMapDiff(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mapDiff(flights)
	}
}

func BenchmarkDoubleFor(b *testing.B) {
	s := map[string]bool{}
	e := map[string]bool{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doubleFor(flights, s, e)
	}
}

func BenchmarkDoubleForCh(b *testing.B) {
	s := sync.Map{}
	e := sync.Map{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doubleForCh(flights, &s, &e)
	}
}

func BenchmarkSliceDiffOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sliceDiff(oneFlight)
	}
}

func BenchmarkMapDiffOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mapDiff(oneFlight)
	}
}

func BenchmarkDoubleForOne(b *testing.B) {
	s := map[string]bool{}
	e := map[string]bool{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doubleFor(oneFlight, s, e)
	}
}

func BenchmarkDoubleForChOne(b *testing.B) {
	s := sync.Map{}
	e := sync.Map{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doubleForCh(oneFlight, &s, &e)
	}
}
