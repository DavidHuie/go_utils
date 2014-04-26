package go_utils

import (
	"math/rand"
)

// Sorts a slice of integers and places the result in the input channel
func quickSortHelper(list []int, c chan []int) {
	length := len(list)

	// Handle base cases
	if length <= 1 {
		c <- list
		return
	}

	// Randomly choose a pivot element
	pivot := int(rand.Int31()) % length

	min_list := make([]int, 0)
	max_list := make([]int, 0)

	for i := 0; i < length; i++ {
		if i == pivot {
			continue
		}

		if list[i] <= list[pivot] {
			min_list = append(min_list, list[i])

		} else {
			max_list = append(max_list, list[i])
		}
	}

	iterations := 0

	min_chan := make(chan []int)
	max_chan := make(chan []int)

	final_list := []int{list[pivot]}

	// Divide and conquer in parallel
	if len(min_list) > 0 {
		iterations += 1
		go quickSortHelper(min_list, min_chan)
	}
	if len(max_list) > 0 {
		iterations += 1
		go quickSortHelper(max_list, max_chan)
	}

	for i := 0; i < iterations; i++ {
		select {
		case l := <-min_chan:
			final_list = append(l, final_list...)
		case l := <-max_chan:
			final_list = append(final_list, l...)
		}
	}

	c <- final_list
}

// Sorts a slice of integers
func QuickSort(list []int) []int {
	c := make(chan []int)
	go quickSortHelper(list, c)
	return <-c
}
