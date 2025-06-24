def max_meetings(start, finish):
    n = len(start)
    # Pair each meeting with (index, start, finish)
    meetings = [(i+1, start[i], finish[i]) for i in range(n)]
    
    # Sort meetings by finish time
    meetings.sort(key=lambda x: x[2])  # Sort by finish

    selected = []
    end_time = 0

    for idx, s, f in meetings:
        if s >= end_time:
            selected.append(idx)
            end_time = f

    return selected

# Test input
s = [1, 3, 0, 5, 8, 5]
f = [2, 4, 6, 7, 9, 9]

result = max_meetings(s, f)
print("Meetings selected:", result)
