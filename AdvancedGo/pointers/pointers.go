package main

type Foo struct {
	Field1 string
	Field2 int
}

// Dont do this
func MakeFoo(f *Foo) error {
	f.Field1 = "val"
	f.Field2 = 30
	return nil
}

// do this
func MakeFoo2() (Foo, error) {
	f := Foo{
		Field1: "val",
		Field2: 30,
	}
	return f, nil
}

func main() {

}
