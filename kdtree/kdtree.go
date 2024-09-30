package kdtree

import (
	"fmt"
	"math"
)

// Point represents a point in multi-dimensional space.
type Point []float64

// Entry хранит ключ (точка) и значение.
type Entry[V any] struct {
	Key   Point
	Value V
}

// KDTree represents a k-d tree data structure.
type KDTree[V any] struct {
	root *Node[V]
	k    int
}

// Node represents a node in the k-d tree.
type Node[V any] struct {
	entry    *Entry[V] // Храним Entry
	left     *Node[V]
	right    *Node[V]
	splitDim int // Dimension used for splitting
}

// NewKDTree creates a new k-d tree with the specified dimensionality.
func NewKDTree[V any](k int) *KDTree[V] {
	return &KDTree[V]{
		root: nil,
		k:    k,
	}
}

// Insert inserts a new point with a value into the k-d tree.
func (tree *KDTree[V]) Insert(key Point, value V) {
	if tree.root == nil {
		tree.root = &Node[V]{
			entry:    &Entry[V]{Key: key, Value: value},
			left:     nil,
			right:    nil,
			splitDim: 0, // Start with the first dimension
		}
		return
	}

	tree.insertRecursive(tree.root, key, value, 0)
}

// insertRecursive recursively inserts a point with a value into the tree.
func (tree *KDTree[V]) insertRecursive(node *Node[V], key Point, value V, depth int) {
	splitDim := depth % tree.k

	// Go left if the point's splitDim coordinate is less than the node's
	if key[splitDim] < node.entry.Key[splitDim] {
		if node.left == nil {
			node.left = &Node[V]{
				entry:    &Entry[V]{Key: key, Value: value},
				left:     nil,
				right:    nil,
				splitDim: (depth + 1) % tree.k,
			}
		} else {
			tree.insertRecursive(node.left, key, value, depth+1)
		}
	} else {
		// Go right otherwise
		if node.right == nil {
			node.right = &Node[V]{
				entry:    &Entry[V]{Key: key, Value: value},
				left:     nil,
				right:    nil,
				splitDim: (depth + 1) % tree.k,
			}
		} else {
			tree.insertRecursive(node.right, key, value, depth+1)
		}
	}
}

// Get finds the entry with the specified key.
func (tree *KDTree[V]) Get(key Point) (*Entry[V], error) {
	node, err := tree.getNodeRecursive(tree.root, key, 0)
	if err != nil {
	    return nil, fmt.Errorf("entry with key %v not found", key)
	}
    return node.entry, nil
    
}

// getNodeRecursive recursively searches for the node with the specified key.
func (tree *KDTree[V]) getNodeRecursive(node *Node[V], key Point, depth int) (*Node[V], error) {
	if node == nil {
		return nil, fmt.Errorf("The root Node can't be nil")
	}

	splitDim := depth % tree.k
	// Go left if the key's splitDim coordinate is less than the node's
	if key[splitDim] < node.entry.Key[splitDim] {
		return tree.getNodeRecursive(node.left, key, depth+1)
	} else if key[splitDim] > node.entry.Key[splitDim] {
		// Go right otherwise
		return tree.getNodeRecursive(node.right, key, depth+1)
	} else {
		// Found the node
		return node, nil
	}
}

// NearestNeighbor finds the nearest neighbor to the given query point.
func (tree *KDTree[V]) NearestNeighbor(query Point) (*Entry[V], float64, error) {
	nearest, minDistance := tree.nearestNeighborRecursive(tree.root, query, 0, math.Inf(1))
	if nearest != nil {
		return nearest.entry, minDistance, nil
	}
	return nil, 0, fmt.Errorf("no nearest neighbor found")
}

// nearestNeighborRecursive recursively searches for the nearest neighbor.
func (tree *KDTree[V]) nearestNeighborRecursive(node *Node[V], query Point, depth int, minDistance float64) (*Node[V], float64) {
	if node == nil {
		return nil, minDistance
	}

	splitDim := depth % tree.k
	distanceToNode := distance(query, node.entry.Key)

	// Check if the current node is closer than the current nearest neighbor
	if distanceToNode < minDistance {
		minDistance = distanceToNode
		nearest := node
		return nearest, minDistance
	}

	// Choose which subtree to explore first based on the query point's coordinate
	var closerSubtree *Node[V]
	if query[splitDim] < node.entry.Key[splitDim] {
		closerSubtree = node.left
	} else {
		closerSubtree = node.right
	}

	// Explore the closer subtree first
	nearest, minDistance := tree.nearestNeighborRecursive(closerSubtree, query, depth+1, minDistance)

	// Explore the further subtree if it might contain a closer point
	if math.Abs(query[splitDim]-node.entry.Key[splitDim]) < minDistance {
		nearest, minDistance = tree.nearestNeighborRecursive(node.right, query, depth+1, minDistance)
	}

	return nearest, minDistance
}

// distance calculates the Euclidean distance between two points.
func distance(p1, p2 Point) float64 {
	sum := 0.0
	for i := 0; i < len(p1); i++ {
		sum += math.Pow(p1[i]-p2[i], 2)
	}
	return math.Sqrt(sum)
}

