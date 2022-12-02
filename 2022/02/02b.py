# A = rock
# B = paper
# C = scissors
#
# X = need to lose
# Y = need to draw
# Z = need to win
#
# 1 point for playing rock
# 2 point for playing paper
# 3 point for playing scissors

total_score = []  # a list to track total score


def prs(opponent, outcome):    # function to play paper rock scissors
    score = 0                  # set score to 0
    if opponent == "A":        # rock
        if outcome == "X":     # you're supposed to lose
            score += 0         # points for lose
            score += 3         # points for scissors
        if outcome == "Y":     # you're supposed to draw
            score += 3         # points for draw
            score += 1         # points for rock
        if outcome == "Z":     # you're supposed to win
            score += 6         # points for win
            score += 2         # points for paper
    if opponent == "B":        # paper
        if outcome == "X":     # you're supposed to lose
            score += 0         # points for lose
            score += 1         # points for rock
        if outcome == "Y":     # you're supposed to draw
            score += 3         # points for draw
            score += 2         # points for paper
        if outcome == "Z":     # you're supposed to win
            score += 6         # points for win
            score += 3         # points for scissors
    if opponent == "C":        # scissor
        if outcome == "X":     # you're supposed to lose
            score += 0         # points for lose
            score += 2         # points for paper
        if outcome == "Y":     # you're supposed to draw
            score += 3         # points for draw
            score += 3         # points for scissors
        if outcome == "Z":     # you're supposed to win
            score += 6         # points for win
            score += 1         # points for rock
    total_score.append(score)  # add round's score to list


with open("../input/02.txt", "r") as file:  # open file
    for line in file:                       # loop through lines
        line = line.split()                 # split items in line
        prs(line[0], line[1])               # send values to function

print(sum(total_score))  # print sum of total_score list
