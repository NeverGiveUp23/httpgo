package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

// using struct to simulate named params

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int64
}

func MyFunc(opts MyFuncOpts) error {
	return nil
}

func AddTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))

	for _, v := range vals {
		out = append(out, base+v)
	}

	return out
}

func Format(format string, values ...interface{}) string {
	return fmt.Sprintf(format, values...)
}

// you can return the names with multiple values
func divAndRemainder(num, denom int) (result, remainder int, err error) {
	if denom == 0 {
		err = errors.New("can't divide by zero")
		return result, remainder, err
	}
	result, remainder = num/denom, num%denom
	return result, remainder, err
}

func add(i, j int) int { return i + j }
func sub(i, j int) int { return i - j }
func mul(i, j int) int { return i * j }
func div(i, j int) int { return i / j }

var opMap = map[string]func(int, int) int{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

// closure returning another function
func makeBase(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func getFile(name string) (*os.File, func(), error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}
	return file, func() {
		if err := file.Close(); err != nil {
			return
		}
	}, nil
}

func main() {
	fmt.Println(AddTo(3))
	fmt.Println(AddTo(2))
	fmt.Println(AddTo(2, 4, 5, 67, 9))

	a := []int{4, 3}

	fmt.Println(AddTo(3, a...))
	fmt.Println(AddTo(3, []int{1, 2, 3, 4, 5}...))

	message := Format("Name: %s, Age: %d, Height: %.2f", "Alice", 30, 5.11)
	fmt.Println(message)

	x, y, _ := divAndRemainder(10, 5)

	fmt.Println("Results:", x, "Remainder:", y)

	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"two", "+", "three"},
		{"5"},
	}

	for _, exp := range expressions {
		if len(exp) != 3 {
			fmt.Println("invalid expression", exp)
			continue
		}
		p1, err := strconv.Atoi(exp[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := exp[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Print("unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(exp[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := opFunc(p1, p2)
		fmt.Println(result)
	}

	twoBase := makeBase(3)
	threeBase := makeBase(4)
	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}

	_, closer, err := getFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	// go does not allow you to unuse variables , so using defer will be best here to ensure you close the function
	defer closer()
}
