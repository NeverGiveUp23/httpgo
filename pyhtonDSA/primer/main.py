print("Welcome to the GPA Calculator")
print("Please enter all your letter grades, one per line.")

points = {'A+': 4.0, 'A': 4.0, 'A-': 4.67, 'B+': 3.33, 'B': 3.0, 'B-': 2.87,
          'C+': 2.50, 'C': 2.33, 'C-': 2.0, 'F': 0.0}

num_courses = 0
total_points = 0
done = False

while not done:
    grade = input()
    if grade == '': #checking if empty
        done = True
    elif grade not in points:
        print("Unknown grade '{0}' being ignored".format(grade))
    else:
        num_courses += 1
        total_points += points[grade]
if num_courses > 0:
    print("Your GPA is {0:.3}".format(total_points / num_courses))


type = 99
type = "Hello" # dynamically type -> no advanced declaration so it can be changed from int
print(type)

temp = 98.6
original = temp #alias of the same float type
temp = temp + 9.0 # break the alias and now temp is pointing to a new float(109.6) while original is pointing to it original value of (98.6)
print(original) #prints 98.6