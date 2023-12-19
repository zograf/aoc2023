# Needed help with this one, I was able to calculate perimiter
# by myself but didn't know how to get area

# I finished part 1 with a grid + floodfill solution, but learned a lot
# in part 2

with open("18.txt", "r") as f:
    lines = f.readlines()
    i = [line.strip().split(" ")[2][2:-1] for line in lines]
    steps = [(x[-1], int(x[:-1], 16)) for x in i]

d = {'0': (0, 1), '1': (1, 0), '2': (0, -1), '3': (-1, 0)}

i, j, a, p = 0, 0, 0, 0
for step in steps:
    direction, length = step
    dj, di = d[direction]
    i, j = i+di*length, j+dj*length
    p, a = p+length, a+i*dj*length
print(a+p//2+1)
