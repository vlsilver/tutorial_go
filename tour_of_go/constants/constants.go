package constants

import "fmt"

// Constants are declared like variables, but with the const keyword.
// Constants can be character, string, boolean, or numeric values.
// Constants cannot be declared using the := syntax.
const pi = 3.14

func constantsDeclare() {
	const World = "world"
	fmt.Printf("Constant Declare: %v\n", pi)
	fmt.Printf("Constant Declare: %v\n", World)
}

// constantsNumeric Numeric constants ar high-precision values.
// An untyped constant takes the type needed by its context.
const (
	// Big Create a huge by number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Small Shift it right again 99 places, so we end up with 1 << 1, or 2
	Small = Big >> 99
)

func constantsNumeric() {
	fmt.Printf("Constants numeric: %v\n", needInt(Small))
	fmt.Printf("Constants numeric: %v\n", needFloat(Small))
	fmt.Printf("Constants numeric: %v\n", needFloat(Big))
}

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x + 0.1
}

func Main() {
	constantsDeclare()
	constantsNumeric()

}
