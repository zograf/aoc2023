from heapq import heappush, heappop

min_steps = 4
max_steps = 10

with open("17.txt", "r") as f:
    lines = f.readlines()
    grid = [list(int(j) for j in list(i.strip())) for i in lines]

seen = set()
q = [(0, 0, 0, 0, 0, 0)]
directions = ((1, 0), (-1, 0), (0, 1), (0, -1))

def push(heat, i, j, dir_i, dir_j, steps):
    new_i, new_j = i + dir_i, j + dir_j
    if new_i >= 0 and new_j >= 0 and new_i < len(grid) and new_j < len(grid[0]):
        heappush(q, (heat+grid[new_i][new_j], new_i, new_j, dir_i, dir_j, steps))

def is_end(i, j, steps):
    return i == len(grid)-1 and j == len(grid[0])-1 and steps >= min_steps

def process(heat, i, j, dir_i, dir_j, steps):
    if steps < max_steps and (dir_i, dir_j) != (0, 0):
        push(heat, i, j, dir_i, dir_j, steps+1)
    if steps >= min_steps or (dir_i, dir_j) == (0, 0):
        for di, dj in directions:
            if (di, dj) != (dir_i, dir_j) and (di, dj) != (-dir_i, -dir_j):
                push(heat, i, j, di, dj, 1)


while len(q):
    heat, i, j, dir_i, dir_j, steps = heappop(q)

    if is_end(i, j, steps):
        print(heat)
        break

    if (i, j, dir_i, dir_j, steps) in seen:
        continue
    seen.add((i, j, dir_i, dir_j, steps))

    
    process(heat, i, j, dir_i, dir_j, steps)
