package variables

import "fmt"

// var1 The `var` statement declares a list of variables; as in function argument lists
// the type is last
// The `var` statement can be at package or function level.
// A var declaration can include initializers, one per variable.
var c, python, java bool

func var1() {
	var i int
	fmt.Printf("Variables: %v, %v, %v, %v \n", i, c, python, java)
}

// var2 with initializers
// A var declaration can include initializers
// one per variable.
// If an initializer is present, the type can be omitted;
// the variable will take the type of the initializer.
var i, j int = 1, 2

func var2() {
	var c, python, java = true, false, "no!"
	fmt.Printf("Variables: %v, %v, %v, %v, %v \n", i, j, c, python, java)
}

// var3 Short variable declarations
// Inside a function, the := short assignment statement ca be used in place
// of a `var` declaration with implicit type.
// Outside a function, every statement begins with
// a keyword(var, func, and so on) and so the := construct is not available.
func var3() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Printf("Variables: %v, %v, %v, %v, %v, %v \n", i, j, c, python, java, k)
}

func Main() {
	var1()
	var2()
	var3()
}
