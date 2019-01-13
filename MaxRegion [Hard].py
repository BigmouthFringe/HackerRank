import math
import os
import random
import re
import sys

def countRegion(grid, i, j):
    count = 1
    grid[i][j] = 0
    
    n, m = len(grid), len(grid[0])
    bounds = [(i - 1, j), (i - 1, j - 1), (i, j - 1), (i + 1, j - 1),
              (i + 1, j), (i + 1, j + 1), (i, j + 1), (i - 1, j + 1)]
    
    for b in bounds:
        if b[0] in range(n) and b[1] in range(m):
            if grid[b[0]][b[1]] == 1 :
                count += countRegion(grid, b[0], b[1])
                       
    return count

def maxRegion(grid):
    max_region = 0   
    n, m = len(grid), len(grid[0])
    
    for i in range(n):
        for j in range(m):
            if grid[i][j] == 1 :
                max_region = max(max_region, countRegion(grid, i, j))
            
    return max_region
                
if __name__ == '__main__':
    n = int(input())
    m = int(input())

    grid = []
    for _ in range(n):
        grid.append(list(map(int, input().rstrip().split())))

    res = maxRegion(grid)

    print(str(res) + '\n')
