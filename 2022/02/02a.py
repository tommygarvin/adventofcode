# A = rock
# B = paper
# C = scissors
#
# X = rock
# Y = paper
# Z = scissors
#
# 1 point for playing rock
# 2 point for playing paper
# 3 point for playing scissors

total_score = []  # a list to track total score


def prs(opponent, you):        # function to play paper rock scissors
    score = 0                  # set score to 0
    if you == "X":             # rock
        score += 1             # points for playing rock
        if opponent == "A":    # rock
            score += 3         # points for draw
        if opponent == "B":    # paper
            score += 0         # points for lose
        if opponent == "C":    # scissors
            score += 6         # points for win
    if you == "Y":             # paper
        score += 2             # points for playing paper
        if opponent == "A":    # rock
            score += 6         # points for win
        if opponent == "B":    # paper
            score += 3         # points for draw
        if opponent == "C":    # scissors
            score += 0         # points for lose
    if you == "Z":             # scissors
        score += 3             # points for playing scissors
        if opponent == "A":    # rock
            score += 0         # points for lose
        if opponent == "B":    # paper
            score += 6         # points for win
        if opponent == "C":    # scissors
            score += 3         # points for draw
    total_score.append(score)  # add round's score to list


with open("../input/02.txt", "r") as file:  # open file
    for line in file:                       # loop through lines
        line = line.split()                 # spplit items in line
        prs(line[0], line[1])               # send values to function

print(sum(total_score))  # print sum of total_score list
