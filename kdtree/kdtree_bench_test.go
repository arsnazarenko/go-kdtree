package kdtree

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

// BenchmarkKDTree_Insert measures the time taken for inserting a large number of points.

func createAndFill(num int)  *KDTree[string] {
	tree := NewKDTree[string](2)
	minV := float64(0)
	maxV := float64(num)
    
    for i := 0; i < num; i++ {
	    rx := minV + rand.Float64()*(maxV-minV)
	    ry := minV + rand.Float64()*(maxV-minV)
		tree.Insert([]float64{float64(rx), float64(ry)}, fmt.Sprintf("Value %d", i))
    }
    return tree
}



func benchmarkKDTree_InsertWithSize(b *testing.B, num int) {
    tree := createAndFill(num)

	minV := float64(0)
	maxV := float64(num)

	rx := minV + rand.Float64()*(maxV-minV)
	ry := minV + rand.Float64()*(maxV-minV)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Insert([]float64{float64(rx), float64(ry)}, fmt.Sprintf("Value %d", i))
	}
}

func BenchmarkKDTree_Insert(b *testing.B) {
    sizes := []int {10, 100, 10000}
    for _, s := range sizes {
        b.Run(fmt.Sprintf("%d", s), func(b *testing.B) {
            benchmarkKDTree_InsertWithSize(b, s)
        })
    }
}

// BenchmarkKDTree_NearestNeighbor measures the time taken for finding the nearest neighbor.
func benchmarkKDTree_NearestNeighborWithSize(b *testing.B, num int) {
    tree := createAndFill(num)
	minV := float64(0)
	maxV := float64(num)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
	    rx := minV + rand.Float64()*(maxV-minV)
	    ry := minV + rand.Float64()*(maxV-minV)
		tree.NearestNeighbor([]float64{rx, ry})
	}
}

func BenchmarkKDTree_NearestNeighbor(b *testing.B) {
    sizes := []int {10, 100, 10000}
    for _, s := range sizes {
        b.Run(fmt.Sprintf("%d", s), func(b *testing.B) {
            benchmarkKDTree_NearestNeighborWithSize(b, s)
        })
    }
}

func benchmarkKDTree_GetWithSize(b *testing.B, num int) {
    tree := createAndFill(num)
	minV := float64(0)
	maxV := float64(num)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
	    rx := minV + rand.Float64()*(maxV-minV)
	    ry := minV + rand.Float64()*(maxV-minV)
		tree.Get([]float64{rx, ry})
	}
}

func BenchmarkKDTree_Get(b *testing.B) {
    sizes := []int {10, 100, 10000}
    for _, s := range sizes {
        b.Run(fmt.Sprintf("%d", s), func(b *testing.B) {
            benchmarkKDTree_GetWithSize(b, s)
        })
    }
}

