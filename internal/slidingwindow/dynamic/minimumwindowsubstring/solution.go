package minimumwindowsubstring

import (
	"maps"
)

/*
Problem

Given two strings s and t, return the *minimum window* substring of s
such that every character in t (including duplicates) is included in the window.
If no such substring exists, return "".

Examples:

Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"

Input: s = "a", t = "aa"
Output: ""   // s doesn't have two 'a's, so no valid window exists

Input: s = "a", t = "a"
Output: "a"

Constraints:

1 <= len(s), len(t) <= 10^5
s and t consist of uppercase and lowercase English letters
Duplicates in t matter — if t = "aab", your window needs at least two as and one b

*/

func minWindow(s, t string) string {
	rsT := []rune(t)
	rsS := []rune(s)

	// LETS MAKE A MAP OF t for source of truth
	var canonicalTMap = make(map[rune]int, len(rsT))
	for _, ru := range rsT {
		canonicalTMap[ru]++
	}

	// keep a clone map to track curr window
	currTMap := maps.Clone(canonicalTMap)

	var (
		min   []rune
		l     int
		found bool
	)
	for r, ru := range rsS {
		if _, ok := canonicalTMap[ru]; ok {
			currTMap[ru]--
			if currTMap[ru] == 0 {
				delete(currTMap, ru)
			}

			if isCurrTMapValid(currTMap) {
				if curr := rsS[l : r+1]; !found {
					min = curr
					found = true
				} else {
					if len(curr) < len(min) {
						min = curr
					}
				}
			}
		}

		for isCurrTMapValid(currTMap) && r-l+1 >= len(rsT) {
			if curr := rsS[l : r+1]; len(curr) < len(min) {
				min = curr
			}

			// try and shrink from the left
			lRu := rsS[l]
			if _, ok := canonicalTMap[lRu]; ok {
				currTMap[lRu]++
			}

			l++
		}
	}

	return string(min)
}

func isCurrTMapValid(tMap map[rune]int) bool {
	for _, v := range tMap {
		if v > 0 {
			return false
		}
	}

	return true
}

// full disclosure i got the below solution from Claude

func claudeMinWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}

	// we need to count the frequency of each character in t
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	required := len(need) // distinct chars that must be fully satisfied

	windowCounts := make(map[byte]int)
	formed := 0 // how many distinct chars currently meet their required count

	l := 0
	bestLen := -1
	bestL := 0

	for r := 0; r < len(s); r++ {
		c := s[r]
		windowCounts[c]++

		// crossed from "not enough" to "exactly enough" for this char
		if need[c] > 0 && windowCounts[c] == need[c] {
			formed++
		}

		// window valid: shrink as far as possible, recording the best
		for formed == required {
			if bestLen == -1 || r-l+1 < bestLen {
				bestLen = r - l + 1
				bestL = l
			}

			left := s[l]
			windowCounts[left]--
			// crossed from "enough" to "no longer enough"
			if need[left] > 0 && windowCounts[left] < need[left] {
				formed--
			}

			l++
		}
	}

	if bestLen == -1 {
		return ""
	}
	return s[bestL : bestL+bestLen]
}
