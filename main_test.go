package main

import (
	"runtime"
	"testing"
)

func BenchmarkFeatureSelection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		run(1)
	}

}

func BenchmarkFeatureSelectionParallell(b *testing.B) {
	for i := 0; i < b.N; i++ {
		run(runtime.NumCPU())
	}

}
