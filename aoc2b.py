x = 0
y = 0

with open('input2a.txt') as file:
    aim = 0

    for i in file:
        isp = i.split(' ')
        if isp[0] == 'forward':
            x += int(isp[1])
            y += int(isp[1]) * aim
        if isp[0] == 'down':
            aim += int(isp[1])
        if isp[0] == 'up':
            aim -= int(isp[1])

print(x*y)