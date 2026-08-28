package maxsumsubarrayofsizek

/*
Problem: Maximum Sum Subarray of Size K

Given an array of positive integers and an integer k, find the maximum sum of any contiguous subarray of size k.

Example:

Input: nums = []int{2, 1, 5, 1, 3, 2}, k = 3
Output: 9

(The subarray [5, 1, 3] has the max sum of 9.)
*/

func maximumSubarraySum(nums []int, k int) int {
	// establish window
	var windowSum int = 0
	for i := range k {
		windowSum += nums[i]
	}

	// slide window until the end
	maxSum := windowSum
	for i := k; i < len(nums); i++ {
		windowSum = windowSum - nums[i-k] + nums[i]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return maxSum
}
