import math
import os
import random
import re
import sys

def whatFlavors(cost, money):
    n = len(cost)
    hashMap = {}
    
    for i in range(0, n - 1):
        y = money - cost[i]
        if y in hashMap :
            return [i + 1, cost.index(y) + 1]
        hashMap[cost[i]] = cost[i]
    
if __name__ == '__main__':
    t = int(input())

    for t_itr in range(t):
        money = int(input())
        n = int(input())

        cost = list(map(int, input().rstrip().split()))

        result = whatFlavors(cost, money)
        result.sort()
        print(result[0], result[1])
