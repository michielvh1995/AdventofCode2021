incs = 0

with open('input.txt') as fil:
    # Elke stap gooi je de oudste uit je window.
    arr = []
    for i in fil:
        arr.append(int(i))
    
    pv = sum(arr[:3])

    for i in range(1, len(arr)-2):
        lv = sum(arr[i:i+3])
        if pv < lv:
            incs += 1
        pv = lv
        

print(incs)