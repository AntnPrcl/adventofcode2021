

file = open("/home/apourcel/Documents/Scripts/adventofcode2021/inputs/day04.txt", "r")

nums = file.readline()
nums = nums.replace("\n", "").split(",")

for num in nums:
    num = int(num)
print(nums)
file.readline()
input = file.read()
input = input.split("\n\n")
tables=[]
i = 0
j = 0
k = 0

for table in input:
    table = table.split("\n")
    new_table = []
    for row in table:
        row = row.replace("  ", " ").split(" ")
        if row[0] == '':
            row.pop(0)
        new_row = []
        for column in row:
            new_row.append([int(column), 0])
        new_table.append(new_row)
    tables.append(new_table)

print(tables)






