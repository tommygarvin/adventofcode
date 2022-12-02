elf_totals = []                                 # create list to track totals

with open("../input/01.txt", "r") as file:      # open file
    elves = file.read().split("\n\n")           # split by empty newline
    elf_index = 0                               # track elf index
    for elf in elves:                           # loop through elves
        backpack = elf.split("\n")              # put the elf's items in a list
        backpack = [int(i) for i in backpack]   # convert strings to int
        elf_totals.append(sum(backpack))        # add results to list
elf_sorted = sorted(elf_totals)                 # sort the list

print(elf_sorted[-1])                           # get last item in list
print(sum(elf_sorted[-3:]))                     # sum of the last 3 items
