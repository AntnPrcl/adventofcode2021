
import numpy

file = open("/home/apourcel/Documents/Scripts/adventofcode2021/inputs/day04.txt", "r")

nums = file.readline()
nums = nums.replace("\n", "").split(",")

for num in nums:
    num = int(num)
print(nums)
file.readline()
input = file.read()
input = input.split("\n\n")
tables=numpy(len(input), 5, 5, 2)
i = 0
j = 0
k = 0
tables[0][0][0] = [12, 0]

for i, table in enumerate(input):
    table = table.split("\n")
    for j, row in enumerate(table):
        row = row.replace("  ", " ").split(" ")
        if row[0] == '':
            row.pop(0)
        for k, column in enumerate(row):
            tables[i][j][k] = [int(column), 0]
            print("")

print(tables)






