def findSquares(x,y):
    squares = 0
    for i in range(0, min(x, y)):
        squares += x * y
        x -= 1
        y -= 1
    return squares