package main

import (
	"math"
	"math/cmplx"
)

// Variables
var globalVar1, globalVar2, globalVar3 bool

// Basic types
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// Functions continued
func add(x, y int) int {
	return x + y
}

// Multiple results
func swap(x, y string) (string, string) {
	return y, x
}

// Named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	// fmt.Println(math.Pi)
	// Functions
	// fmt.Println(add(1, 2))

	// Multiple results
	// a, b := swap("world", "hello")
	// fmt.Println(a, b)

	// Named return values
	// fmt.Println(split(17))

	// Short variable declarations
	// Outside a function, := construct is not available
	// lk := 3
	// lc, lpython, ljava := true, false, "no!"
	// fmt.Println(lk, lc, lpython, ljava)

	// Basic types
	// fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	// fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	// fmt.Printf("Type: %T Value: %v\n", z, z)

	// Zero values
	// var (
	// 	i int
	// 	f float64
	// 	b bool
	// 	s string
	// )

	// fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// Type conversions
	// Go assignment between items of different type requires an explicit conversion.
	var x, y int = 3, 4
	var s float64 = math.Sqrt(x)
	var f float64 = math.Sqrt(float64(x + x))
	// var z uint = uint(f)

	// fmt.Println(x, y, f, z)

	// Type inference
	// The variable's type is inferred from the value on the right hand side.
	// v := 0.867 + 0.5i // change me!
	// fmt.Printf("v is of type %T\n", v)

	// Constants cannot be declared using the := syntax.
	// const Pi = 3.14
	// fmt.Println(Pi)

	// Numeric Constants
	// const (
	// 	Big   = 1 << 100
	// 	Small = Big >> 99
	// )

	// Numeric Constants
	// fmt.Println(Small)
	// fmt.Println(needFloat(Small))
	// fmt.Println(needFloat(Big))
}
