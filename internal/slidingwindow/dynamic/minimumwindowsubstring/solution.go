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
	var tMap = make(map[rune]int, len(rsT))
	for _, ru := range rsT {
		i := tMap[ru]
		tMap[ru] = i + 1
	}

	// keep a clone map to track curr window
	currTMap := maps.Clone(tMap)

	var (
		min   []rune
		l     int
		found bool
	)
	for r, ru := range rsS {
		if _, ok := tMap[ru]; ok {
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
			if _, ok := tMap[lRu]; ok {
				i := currTMap[lRu]
				currTMap[lRu] = i + 1
			}

			l++
		}
	}

	if !found {
		return ""
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
