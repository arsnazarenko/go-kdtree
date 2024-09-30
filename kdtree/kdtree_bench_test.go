package kdtree

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

// BenchmarkKDTree_Insert measures the time taken for inserting a large number of points.
func BenchmarkKDTree_Insert(b *testing.B) {
	tree := NewKDTree[string](2)

	minV := float64(0)
	maxV := float64(b.N)

	rx := minV + rand.Float64()*(maxV-minV)
	ry := minV + rand.Float64()*(maxV-minV)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Insert([]float64{float64(rx), float64(ry)}, fmt.Sprintf("Value %d", i))
	}
}

// BenchmarkKDTree_NearestNeighbor measures the time taken for finding the nearest neighbor.
func BenchmarkKDTree_NearestNeighbor(b *testing.B) {
	tree := NewKDTree[string](2)
	minV := float64(0)
	maxV := float64(b.N)

	rx := minV + rand.Float64()*(maxV-minV)
	ry := minV + rand.Float64()*(maxV-minV)

	// Insert a large number of points for benchmarking.
	for i := 0; i < 10000; i++ {
		tree.Insert([]float64{float64(rx), float64(ry)}, fmt.Sprintf("Value %d", i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.NearestNeighbor([]float64{float64(i), float64(i)})
	}
}

// BenchmarkKDTree_GetNode measures the time taken for getting a node by key.
func BenchmarkKDTree_GetNode(b *testing.B) {
	tree := NewKDTree[string](2)
	// Insert a large number of points for benchmarking.
	for i := 0; i < b.N; i++ {
		tree.Insert([]float64{float64(i), float64(i)}, fmt.Sprintf("Value %d", i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Get([]float64{float64(i), float64(i)})
	}
}
