file = open("../inputs/day01.txt", "r")
input=file.read()
input=input.split("\n")

print("*** ADVENT OF CODE ** DAY 01 ***")
print("====== PART I ======")
nums = []
for i in range (len(input)) :
    nums.append(int(input[i]))

cnt = 0
for i in range(0, len(nums)-1, 1):
    if nums[i] < nums[i+1]:
        cnt += 1
print(f'Ans: {cnt}') # 7



print("\n====== PART II ======")
cnt = 0
for i in range(0, len(nums)-3, 1):
    if nums[i] + nums[i+1] + nums[i+2] < nums[i+1] + nums[i+2] + nums[i+3]:
        cnt += 1
print(f'Ans: {cnt}') # 7