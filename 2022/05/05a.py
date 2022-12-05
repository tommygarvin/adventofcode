# nested list representing stacks of crates
stacks = [
    ["None"],
    ["W", "M", "L", "F"],
    ["B", "Z", "V", "M", "F"],
    ["H", "V", "R", "S", "L", "Q"],
    ["F", "S", "V", "Q", "P", "M", "T", "J"],
    ["L", "S", "W"],
    ["F", "V", "P", "M", "R", "J", "W"],
    ["J", "Q", "C", "P", "N", "R", "F"],
    ["V", "H", "P", "S", "Z", "W", "R", "B"],
    ["B", "M", "J", "C", "G", "H", "Z", "W"]
]

# test data
# stacks = [
#     ["None"],
#     ["Z", "N"],
#     ["M", "C", "D"],
#     ["P"]
# ]


def move_crane(pattern):                   # function to move the crane
    count = pattern[0]                     # pull data out of pattern list
    source = pattern[1]                    # pull data out of pattern list
    destination = pattern[2]               # pull data out of pattern list
    for i in range(count):                 # loop number of times in count
        crate = stacks[source].pop()       # pop last crate from source stack
        stacks[destination].append(crate)  # add crate to destnation stack


with open("../input/05.txt", "r") as file:  # read from file
    for line in file:                       # loop through lines in file
        pattern = []                        # an empty list to hold pattern
        for s in line.split():              # split up the line into strings
            if s.isdigit():                 # if the string is a digit
                pattern.append(int(s))      # then add it to the pattern list
        move_crane(pattern)                 # send pattern to crane

result = []                   # an empty list to hold results
for stack in stacks:          # for each nested list
    result.append(stack[-1])  # add last item to results
print(result)                 # print results
