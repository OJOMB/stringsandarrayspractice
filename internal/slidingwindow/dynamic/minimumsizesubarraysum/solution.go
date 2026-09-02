package minimumsizesubarraysum

/*
Problem: Minimum Size Subarray Sum

Given an array of positive integers nums and a positive integer target, find the length of the shortest contiguous subarray whose sum is ≥ target. If no such subarray exists, return 0.

Example:

Input: target = 7, nums = []int{2, 3, 1, 2, 4, 3}
Output: 2

(The subarray [4, 3] has sum 7, and is the shortest one meeting the target)
*/

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	minLen := n + 1 // sentinel value, bigger than any possible valid length
	sum := 0
	left := 0

	for right := range n {
		sum += nums[right] // always grow

		// we only ever shrink the window when the current sum is greater than or equal to the target
		// we only increment right, in this manner we make sure we explore all possible subarrays in O(n) time
		for sum >= target {
			// shrink while still valid
			if right-left+1 < minLen {
				minLen = right - left + 1
			}

			sum -= nums[left]
			left++
		}
	}

	if minLen == n+1 {
		return 0
	}

	return minLen
}
