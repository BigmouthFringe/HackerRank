import math
import os
import random
import re
import sys

# Complete the countInversions function below.
def mergesort(arr):
    temp = [0] * len(arr)
    return _mergesort(arr, temp, 0, len(arr) - 1)

def _mergesort(arr, temp, left_start, right_end):
    count = 0

    if left_start >= right_end : 
        return count
    middle = int((left_start + right_end) / 2)
    count += _mergesort(arr, temp, left_start, middle)
    count += _mergesort(arr, temp, middle + 1, right_end)
    count += _mergeHalves(arr, temp, left_start, right_end)

    return count

def _mergeHalves(arr, temp, left_start, right_end):
    count = 0

    left_end = int((right_end + left_start) / 2)
    right_start = left_end + 1
    size = right_end - left_start + 1

    left = left_start
    right = right_start
    index = left_start

    while left <= left_end and right <= right_end :
        if arr[left] <= arr[right] :
            temp[index] = arr[left]
            left += 1
        else:
            temp[index] = arr[right]
            right += 1
            count += left_end - left + 1
        index += 1

    _arrayCopy(arr, left, temp, index, left_end - left + 1)
    _arrayCopy(arr, right, temp, index, right_end - right + 1)
    _arrayCopy(temp, left_start, arr, left_start, size)

    return count

def _arrayCopy(src, src_pos, dest, dest_pos, len):
    for i in range(len):
        dest[i + dest_pos] = src[i + src_pos]

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    t = int(input())

    for t_itr in range(t):
        n = int(input())
        arr = list(map(int, input().rstrip().split()))
        result = mergesort(arr)
        fptr.write(str(result) + '\n')

    fptr.close()
