package tests

import (
	"leetcode/medianoftwosortedlists"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	arr1 := []int{1, 3}
	arr2 := []int{2}
	result := medianoftwosortedlists.FindMedianSortedArrays(arr1, arr2)
	if result != 2 {
		t.Error("RETURN", result, "SOLUTION", 2)
	}
	arr1 = []int{1, 3, 5, 7}
	arr2 = []int{2, 4, 6}
	result = medianoftwosortedlists.FindMedianSortedArrays(arr1, arr2)
	if result != 4 {
		t.Error("RETURN", result, "SOLUTION", 4)
	}
	arr1 = []int{3, 3, 3, 3}
	arr2 = []int{3, 3, 3, 3, 3}
	result = medianoftwosortedlists.FindMedianSortedArrays(arr1, arr2)
	if result != 3 {
		t.Error("RETURN", result, "SOLUTION", 3)
	}
	arr1 = []int{1, 3, 5, 7}
	arr2 = []int{2, 4, 6, 8}
	result = medianoftwosortedlists.FindMedianSortedArrays(arr1, arr2)
	if result != 4.5 {
		t.Error("RETURN", result, "SOLUTION", 4.5)
	}
	arr1 = []int{1, 2}
	arr2 = []int{3, 4}
	result = medianoftwosortedlists.FindMedianSortedArrays(arr1, arr2)
	if result != 2.5 {
		t.Error("RETURN", result, "SOLUTION", 2.5)
	}
	arr1 = []int{0, 0, 0, 0, 0}
	arr2 = []int{-1, 0, 0, 0, 0, 0, 1}
	result = medianoftwosortedlists.FindMedianSortedArrays(arr1, arr2)
	if result != 0 {
		t.Error("RETURN", result, "SOLUTION", 0)
	}
	arr1 = []int{1, 2}
	arr2 = []int{-1, 3}
	result = medianoftwosortedlists.FindMedianSortedArrays(arr1, arr2)
	if result != 1.5 {
		t.Error("RETURN", result, "SOLUTION", 1.5)
	}
	arr1 = []int{1, 3}
	arr2 = []int{2, 5}
	result = medianoftwosortedlists.FindMedianSortedArrays(arr1, arr2)
	if result != 2.5 {
		t.Error("RETURN", result, "SOLUTION", 2.5)
	}
	arr1 = []int{1, 2, 2}
	arr2 = []int{1, 2, 3}
	result = medianoftwosortedlists.FindMedianSortedArrays(arr1, arr2)
	if result != 2 {
		t.Error("RETURN", result, "SOLUTION", 2)
	}
	arr1 = []int{100001}
	arr2 = []int{100000}
	result = medianoftwosortedlists.FindMedianSortedArrays(arr1, arr2)
	if result != 100000.5 {
		t.Error("RETURN", result, "SOLUTION", 100000.5)
	}
	arr1 = []int{2, 3, 4}
	arr2 = []int{1}
	result = medianoftwosortedlists.FindMedianSortedArrays(arr1, arr2)
	if result != 2.5 {
		t.Error("RETURN", result, "SOLUTION", 2.5)
	}
	arr1 = []int{3}
	arr2 = []int{1, 2, 4}
	result = medianoftwosortedlists.FindMedianSortedArrays(arr1, arr2)
	if result != 2.5 {
		t.Error("RETURN", result, "SOLUTION", 2.5)
	}
	arr1 = []int{1, 2, 4}
	arr2 = []int{3}
	result = medianoftwosortedlists.FindMedianSortedArrays(arr1, arr2)
	if result != 2.5 {
		t.Error("RETURN", result, "SOLUTION", 2.5)
	}
	arr1 = []int{1, 2, 5}
	arr2 = []int{3, 4, 6}
	result = medianoftwosortedlists.FindMedianSortedArrays(arr1, arr2)
	if result != 3.5 {
		t.Error("RETURN", result, "SOLUTION", 3.5)
	}
}
