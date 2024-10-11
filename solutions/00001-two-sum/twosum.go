package twosum

// Time complexity: O(n)
// Space complexity: O(n)
func twoSum(nums []int, target int) []int {
	val := make(map[int]int)

	for i, v := range nums {
		if j, ok := val[v]; ok {
			return []int{j, i}
		}
		val[target-v] = i
	}

	return nil
}
