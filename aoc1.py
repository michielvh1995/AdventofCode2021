incs = 0

with open('input.txt') as fil:
    pd = 1000000
    for i in fil:
        if int(i) > pd:
            incs += 1
        pd = int(i)

print(incs)