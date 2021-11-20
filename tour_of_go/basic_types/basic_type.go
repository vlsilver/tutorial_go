package basic_types

import (
	"fmt"
	"math"
	"math/cmplx"
)

// Go's bacsic types are
// `bool`
// `string`
// `int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr`
// `byte` -> alias for uint8
// `rune` -> alias for int32, represents a Unicode point
// `float32 float64`
// `complex64 complex128`

// The `int, uint` and `uintptr` types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.
// When you need an integer value you should use int unless you have a specific reason to use a sized
// or unsigned integer type.

var (
	toBe   bool       = false
	maxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func type1() {
	fmt.Printf("Basic Type - Type: %T Value: %v\n", toBe, toBe)
	fmt.Printf("Basic Type - Type: %T Value: %v\n", maxInt, maxInt)
	fmt.Printf("Basic Type - Type: %T Value: %v\n", z, z)
}

// type2 Variables declared without an explicit initial value are given their zero value.
// the zero value is:
// 0 for numeric types
// false for the boolean type, and ""(the empty string) for strings.
func type2() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("Basic Type - Type: %v %v %v %q\n", i, f, b, s)
}

//typeConversion The expression T(v) converts the value v to the type T.

func typeConversion() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	fmt.Printf("Basic Type - Type: %v, %v, %v\n", x, y, f)
}

// When declaring a variable without specifying an explicit type(either by using the := syntax
// or var = expression syntax), the variable's type is inferred from the value on the right hand side.
// When the right hand side of the declaration is typed, the new variable is of that same type.
// But when the right hand side contains an untyped numeric constant, the new variable may be a int, float64
// or complex 128 depending on the precision of the constant.
func typeInference() {
	v := 42 // change me!
	fmt.Printf("Basic Type - Type: v is of type %T\n", v)
}

func Main() {
	type1()
	type2()
	typeConversion()
	typeInference()
}
