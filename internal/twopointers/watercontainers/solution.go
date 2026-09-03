package watercontainers

/*
Problem

You're given an integer array height of length n.
There are n vertical lines drawn such that the two endpoints of the i-th line are (i, 0) and (i, height[i]).
Find two lines that, together with the x-axis, form a container that holds the most water. Return the maximum amount of water it can hold.

Note: you can't tilt the container — the amount of water held is determined by the shorter of the two chosen lines, times the distance between them.

Examples:

Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
// lines at index 1 (height 8) and index 8 (height 7)
// width = 8 - 1 = 7, height = min(8,7) = 7
// area = 7 * 7 = 49

Input: height = [1,1]
Output: 1

Constraints:
2 <= n <= 10^5
0 <= height[i] <= 10^4
*/

func maxArea(heights []int) int {
	// left starts from the first and right is starting from the last index point
	l, r := 0, len(heights)-1
	best := 0

	for l < r {
		width := r - l
		heightL := heights[l]
		heightR := heights[r]

		height := heightL
		var lForward = true
		if heightL > heightR {
			height = heightR
			lForward = false
		}
		area := width * height

		if area > best {
			best = area
		}

		// move whichever pointer inwards (default is l if both are same)
		if lForward {
			l++
		} else {
			r--
		}
	}

	return best
}
