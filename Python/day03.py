import collections

file = open("input", "r")
input=file.read().split("\n")

#Part 1
def most_common_bit(index, input):
    nb1=0
    nb0=0
    for i in range(len(input)-1):
        if input[i][index] == '1':
            nb1+=1
        if input[i][index] == '0':
            nb0+=1
    #if nb1 < nb0:
       # print('0')
    #else:
       # print('1')

for i in range(len(input[0])):
    most_common_bit(i, input)
rate = input
#Part 2
gamma = rate[::]
for i in range(len(rate[0])):
    most = collections.Counter([r[i] for r in gamma])
    most = '1' if most['1'] >= most['0'] else '0'
    gamma = list(filter(lambda x: x[i] == most, gamma))
    if len(gamma) == 1:
        break

epsilon = rate[::]
for i in range(len(rate[0])):
    least = collections.Counter([r[i] for r in epsilon])
    least = '0' if least['1'] >= least['0'] else '1'
    epsilon = list(filter(lambda x: x[i] == least, epsilon))
    if len(epsilon) == 1:
        break

epsilon_rate = int(epsilon[0], 2)
gamma_rate = int(gamma[0], 2)
print(f'Ans: {gamma_rate * epsilon_rate}')