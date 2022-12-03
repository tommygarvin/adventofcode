def find_duplicate(compartment1, compartment2):  # function to find duplicate
    duplicate = ""                               # declare string variable
    for i in range(0, len(compartment1)):        # loop through length of string
        if compartment1[i] in compartment2:      # check if character is in string
            duplicate = compartment1[i]          # if yes, capture value
    return duplicate                             # return value


def get_priority(duplicate):            # function to correlate priority
    if duplicate.islower():             # if lowercase
        priority = ord(duplicate) - 96  # adjust unicode result
    if duplicate.isupper():             # if uppercase
        priority = ord(duplicate) - 38  # adjust unicode result
    return priority                     # return value


priorities = []

with open("../input/03.txt", "r") as file:          # open file
    for rucksack in file:                           # loop through rucksacks
        compartment1 = rucksack[:len(rucksack)//2]  # first half
        compartment2 = rucksack[len(rucksack)//2:]  # second half
        priorities.append(get_priority(
            find_duplicate(compartment1, compartment2)))  # send results to priorities list

print(sum(priorities))  # sum the priorties list
