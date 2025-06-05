# While Loops
# break, continue, pass
x = 0

while x < 5:
    print(f'The current value of x is {x}')
    x+= 1
else:
    print('X is not less than 5')

y = [1, 2, 3]
for item in y:
    pass

print('end of my script')

mystring = 'Sammmy'

for letter in mystring:
    if letter == 'm':
        break
    print(letter)
