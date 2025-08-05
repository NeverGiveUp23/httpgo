package main

import "fmt"

// using value and pointer semantics -> when using the value semantics you should look at that struct data type and realize if that data should be changed or not

type Person struct {
	name string
}

// here we take the name data in our Person struct and RETURN the copy and work with the copy, this allows us to keep the original data untouched in memory

func (p Person) ChangeNameKeepOriginal(name string) Person {
	p.name = name
	fmt.Println(p.name, "Inside the function with te copy")
	return p
}

func main() {

	p1 := Person{name: "Felix"}
	fmt.Println(p1.name, "The original value")
	p1 = p1.ChangeNameKeepOriginal("Jeff")
	fmt.Println(p1.name, "The copy change returning")

}
