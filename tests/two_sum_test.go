package tests

import (
	"fmt"
	"leetcode/twosum"
	"testing"
)

//go test tests/two_sum_test.go -v

func TestTwoSum(t *testing.T) {
	nums := [...]int{2, 7, 11, 15}
	target := 9
	values := twosum.TwoSum(nums[:], target)
	fmt.Println(values)
}
