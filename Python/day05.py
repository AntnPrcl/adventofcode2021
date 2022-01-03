MAP_SIZE = 1000




def read_data():
    file = open("inputs/day05.txt")
    input = file.read()
    input = input.split("\n")
    data = []
    for line in input:
        new_line = []
        line = line.split(" -> ")
        for point in line:
            new_line.append(int(point.split(",")[0]))
            new_line.append(int(point.split(",")[1]))
        data.append(new_line)
    return data

def generate_empty_map():
    map = []
    for i in range(MAP_SIZE):
        new_line = []
        for j in range(MAP_SIZE):
            new_line.append(0)
        map.append(new_line)
    return map

def filter_data(old_input: list):
    new_data = []
    for line in old_input:
        if line[0] == line[2] or line[1] == line[3]:
            new_data.append(line)
    return new_data

def count_cell(map: list):
    count = 0
    for line in map:
        for columns in line:
            if columns > 1:
                count += 1
    return count

def write_lines(input : list, map: list):
    for line in input:
        if line[0] == line[2]:
            if line[1] < line[3]:
                for i in range(line[1], line[3]+1, 1):
                    map[i][line[0]] += 1
            else:
                for i in range(line[3], line[1]+1, 1):
                    map[i][line[0]] += 1
        else:
            if line[0] < line [2]:
                for i in range(line[0], line[2]+1, 1):
                    map[line[1]][i] += 1
            else:
                for i in range(line[2], line[0]+1, 1):
                    map[line[1]][i] += 1

    print("Ans : " + str(count_cell(map)))

def write_lines_II(input: list, map : list):
    for line in input:
        if line[0] == line[2] or line[1] == line[3]:
            if line[0] == line[2]:
                if line[1] < line[3]:
                    for i in range(line[1], line[3]+1, 1):
                        map[i][line[0]] += 1
                else:
                    for i in range(line[3], line[1]+1, 1):
                        map[i][line[0]] += 1
            else:
                if line[0] < line [2]:
                    for i in range(line[0], line[2]+1, 1):
                        map[line[1]][i] += 1
                else:
                    for i in range(line[2], line[0]+1, 1):
                        map[line[1]][i] += 1
        else : 
            if line[0] < line[2]:
                line_size = line[2] - line[0] + 1
                if line[1] < line [3]:
                    for i in range(line_size):
                        map[line[1] + i][line[0] + i] += 1
                else:
                    for i in range(line_size):
                        map[line[1] - i][line[0] + i] += 1
            else:
                line_size = line[0] - line[2] + 1
                if line[1] < line [3]:
                    for i in range(line_size):
                            map[line[1] + i][line[0] - i] += 1
                else:
                    for i in range(line_size):
                            map[line[1] - i][line[0] - i] += 1

    print("Ans : " + str(count_cell(map)))

data = read_data()
map = generate_empty_map()
#data = filter_data(data)
#write_lines(data, map)
write_lines_II(data, map)