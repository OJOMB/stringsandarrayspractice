package maxaveragesubarray

/*

643. Maximum Average Subarray I
Easy

You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value.
Any answer with a calculation error less than 10-5 will be accepted.

Example 1:

Input: nums = [1,12,-5,-6,50,3], k = 4
Output: 12.75000
Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75
Example 2:

Input: nums = [5], k = 1
Output: 5.00000

Constraints:

n == nums.length
1 <= k <= n <= 105
-104 <= nums[i] <= 104

*/

func findMaxAverage(nums []int, k int) float64 {
	n := len(nums)
	if n < k {
		return 0
	}

	sum := 0
	for i := range k {
		sum += nums[i]
	}

	// remember we don't care about dividing by k until the end
	maxSum := sum
	for i := k; i < n; i++ {
		// add int at index i and remove int at index i-k, thus maintaining a fixed window of size k
		sum += nums[i] - nums[i-k]
		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}
