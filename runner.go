package main

import (
	"fmt"
	"github.com/go-learn-algorithms/sorting"
)

func main() {
	arr := []int{5,1,4,2,3}

	sorting.MergeSortInitial(arr)
	fmt.Print(arr)
}
