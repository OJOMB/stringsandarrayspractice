package rottenoranges

/*
Problem

You're given an m x n grid where each cell can have one of three values:

0 — empty cell
1 — fresh orange
2 — rotten orange

Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.
Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1.

Examples:

Input:
grid = [
  [2,1,1],
  [1,1,0],
  [0,1,1]
]
Output: 4

Input:
grid = [
  [2,1,1],
  [0,1,1],
  [1,0,1]
]
Output: -1
// the orange in the bottom-left corner is isolated by the 0, can never rot

Input:
grid = [[0,2]]
Output: 0
// no fresh oranges to begin with

Constraints:

1 <= m, n <= 10
grid[i][j] is 0, 1, or 2
*/

func orangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	type vertex struct{ r, c int }
	var queue []vertex
	fresh := 0

	// TODO: scan grid, seed queue with all rotten oranges, count fresh ones
	for r := range rows {
		for c := range cols {
			switch grid[r][c] {
			case 2:
				queue = append(queue, vertex{r, c})
			case 1:
				fresh++
			}
		}
	}

	minutes := 0
	directions := []vertex{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 && fresh > 0 {
		for range len(queue) {
			curr := queue[0]
			queue = queue[1:]

			for _, d := range directions {
				nr, nc := curr.r+d.r, curr.c+d.c

				// bounds check
				if nr < 0 || nc < 0 || nr >= rows || nc >= cols {
					continue
				}

				// if the neighboring cell contains a fresh orange, rot it and add to the queue
				if grid[nr][nc] == 1 {
					grid[nr][nc] = 2
					fresh--

					queue = append(queue, vertex{nr, nc})
				}
			}
		}

		minutes++
	}

	if fresh > 0 {
		return -1
	}

	return minutes
}
