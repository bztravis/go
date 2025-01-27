// go doesn't have classes, but you can define methods on types
// you can only declare methods on types which are defined in the same package as you're writing the method

package main

import (
	"fmt"
	"image"
	"io"
	"math"
	"strings"
	"time"
)

type Vertex struct {
	X, Y float64
}

// methods are just functions with specicial "receiver" argument lists before the function name
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// methods can receive pointers to types (that are defined in the same module I'm assuming), which is required in order for the method to modify the object. Otherwise methods operate on copies of the value
// also helpful for perf ofc
func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

// an interface type is defined as a set of method signatures
type Abser interface {
	Abs() float64
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

// one of the most common interfaces, implemented by the fmt module. Anything that can describe itself as a string
// type Stringer interface {
// 	String() string
// }

// errors are another common interface. functions often return an error value
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	v.Scale(10) // notice how v is not a pointer, but convenience Go calls the method which accepts a pointer receiver for us
	fmt.Println(v)
	fmt.Println(v.Abs())

	var a Abser = &v // ok, since Abs() is defined on *Vertex
	fmt.Println(a)
	// var b Abser = v // not ok, since Abs() is not defined for Vertex
	// fmt.Println(b)

	// 	If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.

	// In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method M in this example.)

	// Note that an interface value that holds a nil concrete value is itself non-nil.

	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()

	var ii I
	describe(ii)
	// ii.M()	// seg faults, nil concrete value and type means that there's no way to determine which method to call

	// type assertions provide access to an interface value's underlying type
	tt := i.(I)     // returns the concrete value, or panics if assertion fails
	fmt.Println(tt) // prints the concrete value

	ttt, ok := i.(I) // doesn't panic and ok is false if assertion fails
	fmt.Println(ttt, ok)

	// we can have type switch statements
	/*
		switch v := i.(type) {
		case T:
			// here v has type T
		case S:
			// here v has type S
		default:
			// no match; here v has the same type as i
		}
	*/

	aa := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(aa, z)

	if err := run(); err != nil {
		fmt.Println(err)
	}

	// readers are another common interface, which is implemented to read streams of data (files, network io, etc.)
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

// the empty interface: can hold values of any type (every type implements at least zero methods)
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
