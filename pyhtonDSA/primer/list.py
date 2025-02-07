
numbers = [1,8,44,10,5,6]
empty_list = [] #empty list

def getSum(numbers: [int]):
    sum = 0
    for i in numbers:
        sum += i
    print("Total sum: ", sum)

print(getSum(numbers))


pop_last = numbers.pop()


print(pop_last) # will print the last elemnt that is popped -> 6

print(max(numbers))

def insertion_Sort(arr):
    n = len(arr)

    if n <= 1:
        return

    for i in range(1, n):
        key = arr[i]
        j = i - 1
        while j >= 0 and key < arr[j]:
            arr[j + 1] = arr[j]
            j -= 1
        arr[j + 1] = key


insertion_Sort(numbers)
print(numbers)