lines = []
with open('../input/day1.txt') as f:
	lines = f.read().splitlines()

# PART 1
prev = 0
increases = 0
for line in lines: 
	if prev == 0:
		prev = int(line)
		continue
	num = int(line)
	if num > prev:
		increases += 1
	prev = num
print(increases)

# PART 2
sums = []
a = 0
b = 0
for line in lines:
	if a == 0:
		a = int(line)
		continue
	if b == 0:
		b = int(line)	
		continue
	c = int(line)
	sums.append(a + b + c)
	a = b
	b = c
prev = 0
increases = 0
for sum in sums:
	if prev == 0:
		prev = sum
		continue
	if sum > prev:
		increases += 1
	prev = sum
print(increases)

