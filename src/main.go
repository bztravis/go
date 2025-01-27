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

	// pointers
	var sumPtr *int = &sum // pointer to an int, get address
	fmt.Println(*sumPtr)   // dereference

	// structures
	type Vertex struct {
		X int
		Y int
	}

	fmt.Println(Vertex{1, 2})
	vertex := Vertex{0, 0}
	fmt.Println(vertex.X)

	// refering fields of a struct pointer doesn't require special syntax
	vp := &vertex
	fmt.Println(vp.Y)

	fmt.Println(tourGoC())

	fmt.Println(Vertex{X: 1}) // Y:0 is implicit

	var msg [2]string
	msg[0] = "Hello"
	msg[1] = "World"
	fmt.Println(msg)
	fmt.Println(msg[0])
	// fmt.Println(msg[-1]) // does not work

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// slices
	var somePrimes []int = primes[1:3] // inclusive, exclusive
	fmt.Println(somePrimes)

	// slices are like references to arrays
	// multiple slices can reference the same elements of the original array
	var someOtherPrimes []int = primes[1:3] // inclusive, exclusive
	someOtherPrimes[0] = -3

	fmt.Println(primes, somePrimes, someOtherPrimes)

	// you can delcare a slice literal, which creates an underlying array and returns a slice that references it
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	// same slicing rules and features as in python
	s = s[1:4]
	fmt.Println(s)
	s = s[:2]
	fmt.Println(s)
	s = s[1:]
	fmt.Println(s)

	// slices have a length and a capacity, capacity is the number of elements in the underlying array starting from the first element of the slice
	var s1 []int // zero value of slice is nil
	fmt.Println(s1, len(s1), cap(s1))

	// dynamically sized slices: use make
	b1 := make([]int, 0, 5) // len(b)=0, cap(b)=5
	fmt.Println(b1)

	// Slices can contain any type, including other slices
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"

	fmt.Println(board)

	// Go provides an append function
	var s2 []int
	printSlice(s2)

	// append works on nil slices.
	s2 = append(s2, 0)
	printSlice(s2)

	// The slice grows as needed.
	s2 = append(s2, 1)
	printSlice(s2)

	// We can add more than one element at a time.
	s2 = append(s2, 2, 3, 4)
	printSlice(s2)

	// the range for loop iterates over arrays, slices, and maps
	for i, v := range s2 {
		fmt.Println(i, v)
	}
	for _, v := range s2 {
		fmt.Println(v)
	}
	for i := range s2 {
		fmt.Println(i)
	}

	// maps: default value of nil, has to be made
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40, -74,
	}
	fmt.Println(m["Bell Labs"])

	// map literals
	m = map[string]Vertex{
		"Bell Labs": Vertex{
			40, -74,
		},
		"Google": { // Vertex typename in the literal is optional as it can be inferred from the map type
			37, -122,
		},
	}

	// remove key
	delete(m, "Google")

	elem, ok := m["Google"] // ok is boolean, false if doesn't exist in map
	fmt.Println(elem, ok)

	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
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

func tourGoC() (i int) {
	defer func() { i++ }()
	return 1
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// functions are first class objects and can be passed around into other functions and variables
func compute(fn func(float64, float64) float64) float64 {

	i := 0

	// closure
	func() {
		fmt.Println("running")
	}()

	go func(i int) {
		fmt.Println("captured value", i)
	}(i) // best practice to pass in the value to be captured as an argument, since it's not guaranteed to be the same value when the goroutine runs. Since 1.22 for loops have new instsances of variables for every iteration, but this is just to avoid this common mistake within for loops

	i++

	return fn(3, 4)
}

// Type generics
// Accepts a slice of any type as long as the types are comparable (can use == and != on it)
// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}
