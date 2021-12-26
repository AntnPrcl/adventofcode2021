
file = open("inputs/day04.txt", "r")

nums = file.readline()
nums = nums.replace("\n", "").split(",")

file.readline()
input = file.read()
input = input.split("\n\n")
tables=[]

def day04_I(input: list):
    print("====== Part I ======")
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

    for num in nums:
        for i, table in enumerate(tables):
            for j, row in enumerate(table):
                for k, column in enumerate(row):
                    if column[0] == int(num):
                        tables[i][j][k][1] = 1
        
        for table_number, table in enumerate(tables):
            for i in range (5):
                if (table[i][0][1] == 1 and table[i][1][1] == 1 and table[i][2][1] == 1 and table[i][3][1] == 1 and table[i][4][1] == 1) or (table[0][i][1] and table[1][i][1] and table[2][i][1] and table[3][i][1] and table[4][i][1]):
                    print("table number " + str(num) + " won")
                    score = 0
                    for row in table:
                        for column in row:
                            if column[1] == 0:
                                score += column[0]
                    score = score * int(num)
                    print("Ans : " + str(score))
                    exit()

def day04_II(input : list):
    print("====== Part II ======")
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

    for num in nums:
        for i, table in enumerate(tables):
            for j, row in enumerate(table):
                for k, column in enumerate(row):
                    if column[0] == int(num):
                        tables[i][j][k][1] = 1
        
        if len(tables) != 1:
            for j, table in enumerate(tables):
                for i in range (5):
                    if (table[i][0][1] == 1 and table[i][1][1] == 1 and table[i][2][1] == 1 and table[i][3][1] == 1 and table[i][4][1] == 1) or (table[0][i][1] == 1 and table[1][i][1] == 1 and table[2][i][1] == 1 and table[3][i][1] == 1 and table[4][i][1] == 1 ):
                        tables.remove(table)
                        break
        else:
            table = tables[0]
            for i in range(5):
                if (table[i][0][1] == 1 and table[i][1][1] == 1 and table[i][2][1] == 1 and table[i][3][1] == 1 and table[i][4][1] == 1) or (table[0][i][1] == 1 and table[1][i][1] == 1 and table[2][i][1] == 1 and table[3][i][1] == 1 and table[4][i][1] == 1 ):
                    score = 0
                    for row in tables[0]:
                        for column in row:
                            if column[1] == 0:
                                score += column[0]
                    score = score * int(num)
                    print("Ans : " + str(score))
                    exit()

day04_II(input)
