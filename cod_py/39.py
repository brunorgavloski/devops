'''
st = 'Print only the words that start with s in this sentence'

x = st.split()

for letter in x:
    first = letter[0][0]
    if first == 's':
         print(letter)


a = range(11)

for n in a:
    if n%2==1:
        print(n)



list_1 = []
count  = 0
while count <= 50: 
    count = count + 1
    if count % 3 == 0:
        list_1.append(count)
        print(list_1)



st = 'Print every work in this sentence that has an enen number of letter'

x = st.split()

#print(x)
count = 0
for words in x:
    for even in words:       
        count = count + 1
        
    if count % 2 == 1:
        print(words)
        print(count)
        count = 0

st = 'Print every work in this sentence that has an enen number of letter'
x = st.split()

for word in x:
    if len(word) % 2 == 1:
        print(f'{word} ->  {len(word)} letras')

'''

