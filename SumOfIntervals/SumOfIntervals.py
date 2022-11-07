def sum_of_intervals(intervals):
    soi = set([])
    for i in intervals:
        soi.update(range(i[0], i[1]))
    return len(soi)