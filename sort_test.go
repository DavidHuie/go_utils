package go_utils

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	listEqual(t, QuickSort([]int{}), []int{})
	listEqual(t, QuickSort([]int{1}), []int{1})
	listEqual(t, QuickSort([]int{2, 1}), []int{1, 2})
	listEqual(t, QuickSort([]int{3, 2, 1}), []int{1, 2, 3})
	listEqual(t, QuickSort([]int{4, 3, 2, 1}), []int{1, 2, 3, 4})
	listEqual(t, QuickSort([]int{1, 2, 3, 4}), []int{1, 2, 3, 4})
	listEqual(t, QuickSort([]int{4, 4, 4, 4}), []int{4, 4, 4, 4})
}

func listEqual(t *testing.T, l1 []int, l2 []int) {
	if len(l1) != len(l2) {
		t.Errorf("Unequal lists: %s, %s", l1, l2)
	}

	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			t.Errorf("Unequal lists: %s, %s", l1, l2)
		}
	}
}
