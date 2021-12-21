file = open("input", "r")
input1=file.read()
input1=input1.split("\n")

for i in range(len(input1)-1):
	input1[i]=input[i].split(" ")


#Part 1
x=0
z=0
for i in range(len(input1)-1):
	if input1[i][0] == 'forward':
		x=x+int(input1[i][1])
	if input1[i][0] == 'down':
		z=z+int(input1[i][1])
	if input1[i][0] == 'up':
		z=z-int(input1[i][1])
print(" X : " + str(x) + " Depth : " + str(z))

#Part 2
aim=0
x=0
z=0
for i in range(len(input1)-1):
	if input1[i][0] == 'forward':
		x=x+int(input1[i][1])
	        z=z+(int(input1[i][1])*aim)
	if input1[i][0] == 'down':
		aim=aim+int(input1[i][1])
	if input1[i][0] == 'up':
		aim=aim-int(input1[i][1])
print("aim : " + str(aim) + " X : " + str(x) + " Depth : " + str(z))

