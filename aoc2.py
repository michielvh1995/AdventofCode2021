x = 0
y = 0

with open('input2a.txt') as file:    
    for i in file:
        isp = i.split(' ')
        if isp[0] == 'forward':
            x += int(isp[1])
        if isp[0] == 'down':
            y += int(isp[1])
        if isp[0] == 'up':
            y -= int(isp[1])

print(x*y)