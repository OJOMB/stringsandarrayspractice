package numberofislands

/*
Problem

Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically (not diagonally).
You may assume all four edges of the grid are surrounded by water.

Examples:

Input:
grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1

Input:
grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3

Constraints:
1 <= m, n <= 300
grid[i][j] is '0' or '1'
The key insight

Scan every cell. Whenever you hit an unvisited '1', that's the start of a brand-new island — increment your count,
then "flood fill" outward from that cell (BFS or DFS, either works)
marking every connected '1' as visited so the outer scan never counts it again.

The trick most people fumble isn't the algorithm shape, it's the bookkeeping:

How do you mark visited? Two common choices: a separate visited [][]bool grid, or mutate grid[i][j] in place (e.g. flip '1' → '0') to avoid extra space. Either is fine — know the tradeoff (mutating input vs. extra memory).
Bounds checking on every recursive/queue step — going off the grid edges is the classic panic source in Go if you're not careful.
DFS (recursion or an explicit stack) vs BFS (queue) — both are O(m·n) time, just different call/space patterns. Recursion is simpler to write but risks a stack blowup on a huge grid (up to 300×300 = 90,000 cells, which is generally fine for Go's stack, but worth being aware of).

*/

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])

	visited := make([][]bool, rows)
	for r := range visited {
		visited[r] = make([]bool, cols)
	}

	var recurs func(r, c int)
	recurs = func(r, c int) {
		if r < 0 || c < 0 || r >= rows || c >= cols || visited[r][c] == true {
			return
		}

		point := grid[r][c]
		if point != '1' {
			return
		}
		visited[r][c] = true

		recurs(r-1, c)
		recurs(r+1, c)
		recurs(r, c-1)
		recurs(r, c+1)
	}

	count := 0
	for r := range rows {
		for c := range cols {
			if grid[r][c] == '1' && !visited[r][c] {
				count++
				recurs(r, c)
			}
		}
	}

	return count
}
