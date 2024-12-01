
def main():

    left = []
    right = []
    with open('day1.txt', "r") as file:
        for line in file:
            pair = line.split() 
            left.append(int(pair[0]))
            right.append(int(pair[1]))
    left.sort()
    right.sort()

    result1 = sum([abs(pair[0] - pair[1]) for pair in zip(left, right)])
    print(f"Part 1 result: {result1}")

    right_map = {}
    for x in right:
        if x not in right_map:
            right_map[x] = 0
        right_map[x]+=1

    result2 = sum([ x * right_map[x] for x in left if x in right_map]) 

    print(f"Part 2 result: {result2}")
        

main()
