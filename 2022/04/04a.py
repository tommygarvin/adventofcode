def convert_to_set(sections):         # function to convert string to set
    sections = sections.split("-")    # split the start and stop points
    start = int(sections[0])          # convert to int
    stop = int(sections[1]) + 1       # convert to int and add 1
    result = set(range(start, stop))  # create set out of range
    return result                     # send result back


def fully_contain(first, second):                         # function fully contains
    if first.issubset(second) or second.issubset(first):  # check for subset
        fully_contain_count.append(1)                     # if true, enter 1
    else:                                                 # otherwise
        fully_contain_count.append(0)                     # enter 0


fully_contain_count = []  # list to hold hits or misses

with open("../input/04.txt", "r") as file:  # open file
    for line in file:                       # loop through lines
        pair = line.split(",")              # separate first and second elf
        first = convert_to_set(pair[0])     # call function for first elf
        second = convert_to_set(pair[1])    # call function for first elf
        fully_contain(first, second)        # call function with both elves

print(fully_contain_count)       # print list for troubleshooting
print(sum(fully_contain_count))  # print answer
