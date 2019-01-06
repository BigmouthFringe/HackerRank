import math
import os
import random
import re
import sys

def hourglass(arr, a, b):
    h = 0
    h += arr[a][b] + arr[a][b+1] + arr[a][b+2]
    h += arr[a+1][b+1]
    h += arr[a+2][b] + arr[a+2][b+1] + arr[a+2][b+2]
    
    return h
    
def hourglassSum(arr):
    m = len(arr)
    n = len(arr[0])
    max_h = -sys.maxsize - 1
    
    for a in range(m-2):
        for b in range(n-2):
            h = hourglass(arr, a, b)
            if (max_h < h): max_h = h
            
    return max_h

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    arr = []
    for _ in range(6):
        arr.append(list(map(int, input().rstrip().split())))

    result = hourglassSum(arr)

    fptr.write(str(result) + '\n')
    fptr.close()
