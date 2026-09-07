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

func orangesRottingBFS(grid [][]int) int {
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
	for len(queue) > 0 && fresh > 0 {
		// we dont iter on the queue, we iter on the length of the queue at minute x
		// if we iter'd on the queue with a regular for i, rottenOrange := range queue {...} we would not attribute minutes correctly
		for range len(queue) {
			// enqueue
			curr := queue[0]
			// dequeue
			queue = queue[1:]

			// check all neighbouring cells that are within bounds of the grid
			for _, d := range []vertex{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				nr, nc := curr.r+d.r, curr.c+d.c

				// bounds check
				if nr < 0 || nc < 0 || nr >= rows || nc >= cols {
					// skip if out of bounds
					continue
				}

				// if the neighboring cell contains a fresh orange, rot it and add to the queue
				if grid[nr][nc] == 1 {
					grid[nr][nc] = 2
					fresh--

					// we have a new rotten to process in the next minute
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

// func orangesRottingDFS(grid [][]int) int {
// 	if len(grid) == 0 || len(grid[0]) == 0 {
// 		return 0
// 	}

// 	rows, cols := len(grid), len(grid[0])
// 	minTime := make([][]int, rows)
// 	for i := range minTime {
// 		minTime[i] = make([]int, cols)
// 		for j := range minTime[i] {
// 			minTime[i][j] = -1 // -1 = not yet reached
// 		}
// 	}

// 	var dfs func(r, c, t int)
// 	dfs = func(r, c, t int) {
// 		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == 0 {
// 			return
// 		}

// 		// only proceed if this path reaches the cell faster than before
// 		if minTime[r][c] != -1 && minTime[r][c] <= t {
// 			return
// 		}

// 		minTime[r][c] = t
// 		dfs(r+1, c, t+1)
// 		dfs(r-1, c, t+1)
// 		dfs(r, c+1, t+1)
// 		dfs(r, c-1, t+1)
// 	}

// 	freshCount := 0
// 	for r := range rows {
// 		for c := range cols {
// 			switch grid[r][c] {
// 			case 1:
// 				freshCount++
// 			case 2:
// 				minTime[r][c] = 0
// 				dfs(r+1, c, 1)
// 				dfs(r-1, c, 1)
// 				dfs(r, c+1, 1)
// 				dfs(r, c-1, 1)
// 			}
// 		}
// 	}

// 	maxTime := 0
// 	rottenReached := 0
// 	for r := range rows {
// 		for c := range cols {
// 			if grid[r][c] == 1 {
// 				if minTime[r][c] == -1 {
// 					return -1 // unreachable fresh orange
// 				}

// 				rottenReached++
// 				if minTime[r][c] > maxTime {
// 					maxTime = minTime[r][c]
// 				}
// 			}
// 		}
// 	}

// 	if rottenReached < freshCount {
// 		return -1
// 	}

// 	return maxTime
// }
