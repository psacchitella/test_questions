package islandFinder

// countIslands counts the number of islands in the given 2D grid.
// 0 = WATER, 1 = LAND
func CountIslands(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	islandCount := 0

	var exploreIsland func(x, y int)
	exploreIsland = func(x, y int) {
		if x < 0 || x >= rows || y < 0 || y >= cols || grid[x][y] == 0 || visited[x][y] {
			return
		}

		visited[x][y] = true

		// Explore in 4 directions (up, down, left, right)
		exploreIsland(x-1, y) // Up
		exploreIsland(x+1, y) // Down
		exploreIsland(x, y-1) // Left
		exploreIsland(x, y+1) // Right
	}

	// Iterate through the grid
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				islandCount++
				exploreIsland(i, j)
			}
		}
	}

	return islandCount
}
