package sorting

import "math"

// MergeSortInitial is my first-pass attempt at merge sort from memory
func MergeSortInitial(arr []int) {
	length := len(arr)
	if length < 2 {
		return
	}
	midpoint := length/2

	left := arr[0:midpoint]
	right := arr[midpoint:]

	MergeSortInitial(left)
	MergeSortInitial(right)

	copy(arr, merge(left, right))
}

func merge(left []int, right []int) []int {
	size := len(left) + len(right)
	slice := make([]int, size)

	// use channels so we can put sentinel values last
	lChan, rChan := make(chan int, len(left)+1), make(chan int, len(right)+1)
	for _, val := range left {
		lChan <- val
	}
	for _, val := range right {
		rChan <- val
	}
	lChan <- math.MaxInt32
	rChan <- math.MaxInt32

	// initialize our values from the channels
	l, r := <-lChan, <-rChan

	for index := 0; index < size; index++ {
		if l <= r {
			slice[index] = l
			l = <- lChan
		} else {
			slice[index] = r
			r = <- rChan
		}
	}

	return slice
}
