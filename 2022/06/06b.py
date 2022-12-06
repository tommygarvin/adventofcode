with open("../input/06.txt", "r") as file:    # read from file
    for line in file:                         # loop through lines in file
        for i in range(0, len(line)):         # for the length of the line
            if len(set(line[i:i+14])) == 14:  # slice 4 chars, set removes dups
                print(i+14)                   # print index after sequence
                break                         # exit loop
