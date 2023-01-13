package main

import (
	"fmt"
	"leetcode/solutions"
)

//3    2

func main() {
	arr1 := []int{1, 3}
	arr2 := []int{2}

	fmt.Println("RETURN", solutions.FindMedianSortedArrays(arr1, arr2), "SOLUTION", 2)
	arr1 = []int{1, 3, 5, 7}
	arr2 = []int{2, 4, 6}
	fmt.Println("RETURN", solutions.FindMedianSortedArrays(arr1, arr2), "SOLUTION", 4)
	arr1 = []int{3, 3, 3, 3}
	arr2 = []int{3, 3, 3, 3, 3}
	fmt.Println("RETURN", solutions.FindMedianSortedArrays(arr1, arr2), "SOLUTION", 3)
	arr1 = []int{1, 3, 5, 7}
	arr2 = []int{2, 4, 6, 8}
	fmt.Println("RETURN", solutions.FindMedianSortedArrays(arr1, arr2), "SOLUTION", 4.5)

	arr1 = []int{1, 2}
	arr2 = []int{3, 4}
	fmt.Println("RETURN", solutions.FindMedianSortedArrays(arr1, arr2), "SOLUTION", 2.5)

	arr1 = []int{0, 0, 0, 0, 0}
	arr2 = []int{-1, 0, 0, 0, 0, 0, 1}
	fmt.Println("RETURN", solutions.FindMedianSortedArrays(arr1, arr2), "SOLUTION", 0)

	arr1 = []int{1, 2}
	arr2 = []int{-1, 3}
	fmt.Println("RETURN", solutions.FindMedianSortedArrays(arr1, arr2), "SOLUTION", 1.5)

	arr1 = []int{1, 3}
	arr2 = []int{2, 5}
	fmt.Println("RETURN", solutions.FindMedianSortedArrays(arr1, arr2), "SOLUTION", 2.5)

	arr1 = []int{1, 2, 2}
	arr2 = []int{1, 2, 3}
	fmt.Println("RETURN", solutions.FindMedianSortedArrays(arr1, arr2), "SOLUTION", 2)

	arr1 = []int{100001}
	arr2 = []int{100000}
	fmt.Println("RETURN", solutions.FindMedianSortedArrays(arr1, arr2), "SOLUTION", 100000.5)

	arr1 = []int{2, 3, 4}
	arr2 = []int{1}
	fmt.Println("RETURN", solutions.FindMedianSortedArrays(arr1, arr2), "SOLUTION", 2.5)

	arr1 = []int{3}
	arr2 = []int{1, 2, 4}
	fmt.Println("RETURN", solutions.FindMedianSortedArrays(arr1, arr2), "SOLUTION", 2.5)

	arr1 = []int{1, 2, 4}
	arr2 = []int{3}
	fmt.Println("RETURN", solutions.FindMedianSortedArrays(arr1, arr2), "SOLUTION", 2.5)

	arr1 = []int{1, 2, 5}
	arr2 = []int{3, 4, 6}
	fmt.Println("RETURN", solutions.FindMedianSortedArrays(arr1, arr2), "SOLUTION", 3.5)
}
