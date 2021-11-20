package functions

import "fmt"

// A function can take zero or more arguments.

// Notice that the type comes after the variable name.

func add(x int, y int) int {
	fmt.Printf("Functions - add: %v\n", x+y)
	return x + y
}

// add2 We shortened `x int, y int` -> `x, y int`
// When two or more consecutive named function
// parameters share a type, you can omit the type from all but the last.
func add2(x, y int) int {
	fmt.Printf("Functions - add: %v\n", x+y)
	return x + y
}

// swap A function can return any number of results.
func swap(x, y string) (string, string) {
	fmt.Printf("Functions - swap: %v - %v\n", x, y)
	return y, x
}

// split Go's return values may be named.
// If so, they are treated as variables defined at the top of the function
// Naked return statements should be used only in short functions.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	fmt.Printf("Functions - swap: %v - %v\n", x, y)
	return
}

func Main() {
	add(1, 2)
	add2(1, 2)
	swap("hello", "world")
	split(10)
}
