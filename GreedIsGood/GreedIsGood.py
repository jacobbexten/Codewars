from collections import Counter

def score(dice):
    points = {1:1000, 2:200, 3:300, 4:400, 5:500, 6:600}
    counts = Counter(dice)
    total = 0
    
    for key, value in counts.items():
        if value >= 3:
            total += points[key] * (value // 3)
        if key == 1:
            total += 100 * (value % 3)
        elif key == 5:
            total += 50 * (value % 3)
        
    return total