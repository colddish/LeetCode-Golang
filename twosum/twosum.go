package twosum

func TwoSum(nums []int, target int) []int {
	intToIndex := make(map[int]int)
	for i, num := range nums {
		if intToIndex[target-num] != 0 {
			ints := [2]int{i, intToIndex[target-num] - 1}
			return ints[:]
		}
		intToIndex[num] = i + 1
	}
	return nil
}
