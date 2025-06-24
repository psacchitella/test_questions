/**
 * Function to count the number of islands in a given 2D grid.
 * 0 = WATER, 1 = LAND
 */
function countIslands(grid) {
    if (!grid || grid.length === 0) return 0;
    
    const rows = grid.length;
    const cols = grid[0].length;
    const visited = Array.from({ length: rows }, () => Array(cols).fill(false));
    let islandCount = 0;

    function exploreIsland(x, y) {
        if (x < 0 || x >= rows || y < 0 || y >= cols || grid[x][y] === 0 || visited[x][y]) {
            return;
        }
        
        visited[x][y] = true;
        exploreIsland(x - 1, y); // Up
        exploreIsland(x + 1, y); // Down
        exploreIsland(x, y - 1); // Left
        exploreIsland(x, y + 1); // Right
    }

    for (let i = 0; i < rows; i++) {
        for (let j = 0; j < cols; j++) {
            if (grid[i][j] === 1 && !visited[i][j]) {
                islandCount++;
                exploreIsland(i, j);
            }
        }
    }

    return islandCount;
}

// Test the function
const grid = [
    [0, 1, 1, 0, 0],
    [1, 0, 1, 0, 1],
    [0, 0, 0, 1, 1],
    [1, 0, 0, 0, 0]
];

console.log(`Island Count: ${countIslands(grid)}`);
