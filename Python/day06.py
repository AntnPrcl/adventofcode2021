def pass_one_day(aquarium: list):
    new_aquarium = [0, 0, 0, 0, 0, 0, 0, 0, 0]
    for i in range(8):
        new_aquarium[i] += aquarium[i+1]
    new_aquarium[8] += aquarium[0]
    new_aquarium[6] += aquarium[0]
    return new_aquarium

def read_input():
    aquarium = [0, 0, 0, 0, 0, 0, 0, 0, 0]
    with open('inputs/day06.txt', 'r') as file:
        list = file.read().split(",")
        for num in list:
            aquarium[int(num)] += 1
            
    return aquarium

def count_fish(aquarium: list):
    count = 0
    for nb in aquarium:
        count += nb
    return str(count)

def main():
    aquarium = read_input()
    print(aquarium)
    for i in range(256):
        aquarium = pass_one_day(aquarium)
    print("Ans : " + count_fish(aquarium))

main()