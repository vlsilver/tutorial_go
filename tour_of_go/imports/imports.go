package imports

import (
	"fmt"
	"math"
)

// This code group the imports into a parenthesized, "factored" import statement.
// You can also write multiple import statements, like:

// import "fmt"
// import "math"

// But it is good style to use the factored import statement.

// In Go, name is exported if it begins with a capital letter.
// When importing a package, you can refer only to its "exported" names.

func Main() {
	fmt.Printf("Import - Now you have %g programs.\n", math.Sqrt(7))
}



