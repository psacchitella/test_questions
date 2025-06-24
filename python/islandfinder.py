"""
Island Finder Module
This module provides a function to count the number of islands in a given 2D grid.
0 = WATER, 1 = LAND
"""

def count_islands(grid):
    """Counts the number of islands in a given 2D grid using DFS."""
    if not grid or not grid[0]:
        return 0

    rows, cols = len(grid), len(grid[0])
    visited = [[False] * cols for _ in range(rows)]
    island_count = 0

    def explore_island(x, y):
        """Recursive function to explore an island using DFS."""
        if x < 0 or x >= rows or y < 0 or y >= cols or grid[x][y] == 0 or visited[x][y]:
            return
        visited[x][y] = True
        # Explore adjacent cells (up, down, left, right)
        explore_island(x - 1, y)
        explore_island(x + 1, y)
        explore_island(x, y - 1)
        explore_island(x, y + 1)

    # Iterate through the grid to count distinct islands
    for i in range(rows):
        for j in range(cols):
            if grid[i][j] == 1 and not visited[i][j]:
                island_count += 1
                explore_island(i, j)

    return island_count
