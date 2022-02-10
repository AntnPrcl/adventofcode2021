from posixpath import split


def read_input():
    with open("inputs/day07.txt", "r") as file:
        input = file.read()
        input = input.split(',')
        for element in range(len(input)):
            input[element] = int(input[element])
    return input

def count_fuel_cost(crabs: list, position: int):
    count = 0
    for crab in crabs:
        if crab < position:
            count += position - crab
        if crab > position:
            count += crab - position
    return count

def count_fuel_cost_2(crabs: list, position: int):
    count = 0
    for crab in crabs:
        if crab < position:
            diff = position - crab
            for i in range(diff+1):
                count += i
        if crab > position:
            diff = crab - position
            for i in range(diff+1):
                count += i
    return count

crabs = read_input()
cheap_cost = 999999999999999999999999999999999
for i in range(max(crabs)):
    fuel_cost = count_fuel_cost(crabs, i)
    if fuel_cost < cheap_cost:
        cheap_cost = fuel_cost
print("Part I")
print("Ans : " + str(cheap_cost))

cheap_cost = 999999999999999999999999999999999
for i in range(max(crabs)):
    fuel_cost = count_fuel_cost_2(crabs, i)
    if fuel_cost < cheap_cost:
        cheap_cost = fuel_cost
print("Part II")
print("Ans : ", cheap_cost)

