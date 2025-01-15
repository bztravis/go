package main

// import block
import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
)

func add(x, y int) int {
	return x + y
}

// return multiple value
func swap(x, y string) (string, string) {
	return y, x
}

// naked return named values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return x, y
}

var c, python, java bool

var i, j int = 1, 2
var k, l = 3, 4 // type can be omitted from a var declaration if an initializer is present

// short variable initialization: no var keyword or type needed (inferred)
// m := 5	// must be done inside a function scope

// basic types
/*
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/

func returnTrue() bool {
	return true
}

func helloWorld() {
	// defers run when the enclosing function returns
	// defers are added in a stack, so they are LIFO
	defer fmt.Println("!")
	defer fmt.Println("world")

	fmt.Println("hello")
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Println(math.Pi) // capitalized name means it is exported

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	// "zero values" are used if no value is provided for a variable declaration
	var zeroed int
	fmt.Println(zeroed)

	// type conversions
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	i := 42
	ii := float64(i)
	iii := uint(ii)
	fmt.Println(iii)

	// When the right hand side of the declaration is typed, the new variable is of that same type:
	var i2 int
	j2 := i2 // j is an int
	fmt.Println(j2)

	// But when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant:
	i3 := 42           // int
	f3 := 3.142        // float64
	g3 := 0.867 + 0.5i // complex128
	fmt.Println(i3, f3, g3)

	// Constants
	// Constants cannot be declared using the := syntax.
	// I believe capitalization of variables inside of function scope doesn't export them, only if they're on the package level are they exported
	const Pi = 3.14
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	// Numeric constants are high-precision **values**.
	// An untyped constant takes the type needed by its context. It sort of collapses to the type on an individual use basis
	fmt.Println(needInt(Small))
	fmt.Printf("Small is of type %T\n", Small)
	fmt.Println(needFloat(Small))
	fmt.Printf("Small is of type %T\n", Small)
	fmt.Println(needFloat(Big))

	// for loops
	sum := 0
	for counter := 0; counter < 10; counter++ {
		sum += counter
	}
	fmt.Println(sum)

	// init and post statements are optional
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// spin loop
	for {
		fmt.Println("spin once")
		break
	}

	// if
	if sum > 1000 {
		fmt.Println("Greater than 1000!")
	}
	// ifs can have a short statement to execute before the condition like for loops
	if sum = 1000; sum > 1000 {
		fmt.Println("should not see this")
	}

	// variables are generally block scoped like in C
	// but you can access anything available inside of an if statement inside of the else block
	if v := math.Pow(float64(x), 2); v < 2 {
		// return v
	} else {
		fmt.Println(v)
	}

	// switch statements don't need to have break statements
	switch os := runtime.GOOS; os {
	// case returnTrue():
	// 	fmt.Println("True")
	// can call a function but has to have the same return type as the other case values, doesn't return boolean
	case "darwin":
		fmt.Println("Mac OS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("Something else")
	}

	// switch without a condition means that cases are boolean expressions
	switch {
	case returnTrue():
		fmt.Println("This always happens")
	}

	helloWorld()
}

// looks like constant and function definitions are hoisted
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	fmt.Printf("x is of type %T\n", x)
	return x * 0.1
}
