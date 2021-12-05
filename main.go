package main

import (
	"fmt"
	"sorting-algos/sorting"
)

func main() {
	fmt.Println("vim-go")

	nums := []int{33, 32, 1, 3, 5, 7, 54, 24, 99}

	// a slice holds pointer to the undelying array
	sorting.QuickSorter().Sort(nums)

	fmt.Println(nums)

}
