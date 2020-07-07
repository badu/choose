/* This file tests our ability to tally combinations. */

package choose_test

import (
	"testing"

	"github.com/spakin/choose"
)

// Test the Choose function's ability to compute 0 choose 0 through 10 choose 10.
func TestChoose(t *testing.T) {
	binom := [][3]int{
		{0, 0, 1}, {1, 0, 1}, {1, 1, 1}, {2, 0, 1}, {2, 1, 2},
		{2, 2, 1}, {3, 0, 1}, {3, 1, 3}, {3, 2, 3}, {3, 3, 1},
		{4, 0, 1}, {4, 1, 4}, {4, 2, 6}, {4, 3, 4}, {4, 4, 1},
		{5, 0, 1}, {5, 1, 5}, {5, 2, 10}, {5, 3, 10}, {5, 4, 5},
		{5, 5, 1}, {6, 0, 1}, {6, 1, 6}, {6, 2, 15}, {6, 3, 20},
		{6, 4, 15}, {6, 5, 6}, {6, 6, 1}, {7, 0, 1}, {7, 1, 7},
		{7, 2, 21}, {7, 3, 35}, {7, 4, 35}, {7, 5, 21}, {7, 6, 7},
		{7, 7, 1}, {8, 0, 1}, {8, 1, 8}, {8, 2, 28}, {8, 3, 56},
		{8, 4, 70}, {8, 5, 56}, {8, 6, 28}, {8, 7, 8}, {8, 8, 1},
		{9, 0, 1}, {9, 1, 9}, {9, 2, 36}, {9, 3, 84}, {9, 4, 126},
		{9, 5, 126}, {9, 6, 84}, {9, 7, 36}, {9, 8, 9}, {9, 9, 1},
		{10, 0, 1}, {10, 1, 10}, {10, 2, 45}, {10, 3, 120},
		{10, 4, 210}, {10, 5, 252}, {10, 6, 210}, {10, 7, 120},
		{10, 8, 45}, {10, 9, 10}, {10, 10, 1},
	}
	for _, good := range binom {
		n, m, correct := good[0], good[1], good[2]
		bc := choose.NumChoices(n, m)
		if bc != correct {
			t.Fatalf("Expected %d choose %d = %d but computed %d",
				n, m, correct, bc)
		}
	}
}
