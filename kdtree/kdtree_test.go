package kdtree

import "testing"

// TestKDTree_Insert verifies the insertion functionality of the k-d tree.
func TestKDTree_Insert(t *testing.T) {

	// Test cases
	testCases := []struct {
		name  string
		key   Point
		value string
	}{
		{"KDTree insert testcase 1", []float64{3, 4}, "Value 1"},
		{"KDTree insert testcase 2", []float64{5, 6}, "Value 2"},
		{"KDTree insert testcase 3", []float64{2, 7}, "Value 3"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			tree := NewKDTree[string](2)
			tree.Insert(testCase.key, testCase.value)
			// Check by get
			entry, err := tree.Get(testCase.key)
			if err != nil {
				t.Errorf("Expected value by key %v not found", testCase.key)
			}

			if !equalPoints(entry.Key, testCase.key) || entry.Value != testCase.value {
				t.Errorf("Expected value to be [%v, %s], got [%v, %s]", testCase.key, testCase.value, entry.Key, entry.Value)
			}

		})
	}
}

// TestKDTree_NearestNeighbor verifies the nearest neighbor search functionality.
func TestKDTree_NearestNeighbor(t *testing.T) {

	// Test cases
	testCases := []struct {
		name     string
		keys     []Point  // Points to insert
		values   []string // Values to insert
		query    Point    // Query point
		expected string   // Expected value
	}{
		{
			"KDTree search nearest neighbor testcase 1",
			[]Point{
				{3, 4},
				{5, 6},
				{2, 7},
			},
			[]string{"Value 1", "Value 2", "Value 3"},
			[]float64{4, 5},
			"Value 1",
		},
		{
			"KDTree search nearest neighbor testcase 2",
			[]Point{
				{1, 2},
				{3, 5},
				{4, 1},
				{7, 8},
			},
			[]string{"Value 1", "Value 2", "Value 3", "Value 4"},
			[]float64{2, 3},
			"Value 1",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			tree := NewKDTree[string](2)
			for i, key := range testCase.keys {
				tree.Insert(key, testCase.values[i])
			}

			entry, _, err := tree.NearestNeighbor(testCase.query)
			if err != nil {
				t.Errorf("Error during nearest neighbor search: %v", err)
			}
			if entry.Value != testCase.expected {
				t.Errorf("Expected value to be %s, got %s", testCase.expected, entry.Value)
			}

		})
	}
}

// TestKDTree_GetNode verifies the GetNode functionality.
func TestKDTree_Get(t *testing.T) {

	// Test cases
	testCases := []struct {
		name     string
		keys     []Point  // Points to insert
		values   []string // Values to insert
		query    Point    // Query point
		expected string   // Expected value
		found    bool     // Expected result of found
	}{
		{
			"KDTree get value by key testcase 1",

			[]Point{
				{3, 4},
				{5, 6},
				{2, 7},
			},
			[]string{"Value 1", "Value 2", "Value 3"},
			[]float64{3, 4},
			"Value 1",
			true,
		},
		{
			"KDTree get value by key testcase 2",
			[]Point{
				{1, 2},
				{3, 5},
				{4, 1},
				{7, 8},
			},
			[]string{"Value 1", "Value 2", "Value 3", "Value 4"},
			[]float64{2, 3},
			"",
			false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			tree := NewKDTree[string](2)
			for i, key := range testCase.keys {
				tree.Insert(key, testCase.values[i])
			}

			entry, err := tree.Get(testCase.query)
			if err != nil && testCase.found {
				t.Errorf("Error during GetNode: %v", err)
			} else if err == nil && !testCase.found {
				t.Errorf("Expected error but got nil")
			}
			if entry != nil && entry.Value != testCase.expected {
				t.Errorf("Expected value to be %s, got %s", testCase.expected, entry.Value)
			}

		})
	}
}

// equalPoints checks if two points are equal.
func equalPoints(p1, p2 Point) bool {
	if len(p1) != len(p2) {
		return false
	}
	for i := range p1 {
		if p1[i] != p2[i] {
			return false
		}
	}
	return true
}
