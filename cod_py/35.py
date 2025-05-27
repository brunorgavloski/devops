#For loops

mylist = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

for num in mylist:
    if num % 2 == 0:
        print(num)
    else:
        print('Odd Numer {}' .format(num))


print('---------------------')
list_sum = 0

for num in mylist:
    list_sum = list_sum + num

print(f'A soma é: ', list_sum)

print('---------------------')

mystring = 'Hello world'

for letter in mystring:
    print(letter)



print('---------------------')


for letter in 'Hello worl':
    print(letter)

print('---------------------')
tup = (1, 2, 3)
for item in tup:
    print(item)

print('---------------------')
d = {'k1': 1, 'k2': 2, 'k3': 3}

for item in d.items():
    print(item)


print('---------------------')
