e = 0
y = 0

with open('input3.txt') as file:
    c = [0] * 12
    fl = 0

    for i in file: # per line 
        for l in range(len(i)-1):
            c[l] += int(i[l])
        fl += 1
    
    e = ['1' if j > fl/2 else '0' for j in c]
    y = ['0' if j > fl/2 else '1' for j in c]

e = int(str.join('', e),2)
y = int(str.join('', y),2)

print(str(e * y))