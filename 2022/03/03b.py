def find_badge(elf1, elf2, elf3):          # function to find duplicate
    elf1 = set(elf1)                       # convert string to set
    elf2 = set(elf2)                       # convert string to set
    elf3 = set(elf3)                       # convert string to set
    badge = elf1.intersection(elf2, elf3)  # find items that exist in all 3
    return badge.pop()                     # pull out item from set


def get_priority(duplicate):            # function to correlate priority
    if duplicate.islower():             # if lowercase
        priority = ord(duplicate) - 96  # adjust unicode result
    if duplicate.isupper():             # if uppercase
        priority = ord(duplicate) - 38  # adjust unicode result
    return priority                     # return value


badge_values = []

with open("../input/03.txt", "r") as file:  # open file
    data = file.read().splitlines()         # create list of strings
    for i in range(0, len(data), 3):        # from 0 to EOF in step of 3
        elf1 = data[i]                      # elf1 is start index
        elf2 = data[i+1]                    # elf2 is start + 1
        elf3 = data[i+2]                    # elf3 is start + 2
        badge_values.append(get_priority(
            find_badge(elf1, elf2, elf3)))  # send results to list

print(sum(badge_values))  # sum the priorties list
