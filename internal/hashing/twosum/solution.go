package twosum

/*

You are given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]

Constraints:

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.

*/

func TwoSum(nums []int, target int) []int {
	for i, numI := range nums {
		for j, numJ := range nums[i+1:] {
			if numI+numJ == target {
				return []int{i, j + i + 1}
			}
		}
	}

	return nil
}

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for i, v := range nums {
		// if we have seen the remainder before we return it
		if j, ok := seen[target-v]; ok {
			return []int{i, j}
		}

		// if not we store it for future iterations
		seen[v] = i
	}

	return nil
}
