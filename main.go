package main

import (
	"fmt"

	"github.com/igorkrosh/grokking-algorithms-solutions-go.git/algorithms/searching"
)

func main() {
	// Declare an slice
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	reqEl := 5

	// Calling binary search from packet searching
	fmt.Println("---Binary search---")
	fmt.Println("Target array:", arr)
	fmt.Println("Required element: ", reqEl)
	fmt.Println("Founded index: ", searching.BinarySearch(arr, reqEl))
	fmt.Println("------------------")
}
