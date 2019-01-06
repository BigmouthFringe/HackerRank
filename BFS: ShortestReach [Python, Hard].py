from collections import defaultdict

class Graph:
    def __init__(self, size):
        self.graph = defaultdict(list)
        
    def addEdge(self, u, v):
        self.graph[u].append(v)
        self.graph[v].append(u)
        
    def getDistances(self, s, n):
        level = 0      
        queue = []
        distances = [-1] * n
        visited = [False] * n 
        
        queue.append(s)
        visited[s] = True
        
        # Counter for the elements on the current level
        count = 1
        while queue:
            count -= 1
            s = queue.pop(0)
            distances[s] = 6 * level
 
            for i in self.graph[s]:
                if visited[i] == False:
                    queue.append(i)
                    visited[i] = True

            if count == 0: 
                level += 1
                count = len(queue)

        return distances

t = int(input())
for i in range(t):
    n, m = [int(value) for value in input().split()]
    graph = Graph(n)
    
    for i in range(m):
        x, y = [int(x) for x in input().split()]
        graph.addEdge(x-1, y-1) 
        
    try: s = int(input())
    except: s = 1
    distances = graph.getDistances(s-1, n)
    del distances[s-1] 
    print(*distances, sep=" ")
