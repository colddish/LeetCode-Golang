package longestsubstringwithoutrepeatchars

/*
Given a string s, find the length of the longest
substring
 without repeating characters.
*/

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func LengthOfLongestSubstring(s string) int {
	maxDistance := 0
	indexPositions := make(map[rune]int)
	start := 0
	for end, char := range s {
		start = max(start, indexPositions[char])
		maxDistance = max(maxDistance, end-start+1)
		indexPositions[char] = end + 1
	}
	return maxDistance
}
