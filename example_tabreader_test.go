package tabreader

import "fmt"

// func main() {
//     s := newScan(2, 4, 2)
//     var a int
//     var b float32
//     var c int32

//     s.Scan("12 2.2 1", &a, &b, &c)
//     fmt.Printf("%d %f %d\n", a, b, c)

//     s.Scan("1      2", &a, &b, &c)
//     fmt.Printf("%d %f %d\n", a, b, c)

//     s.Scan("        ", &a, &b, &c)
//     fmt.Printf("%d %f %d\n", a, b, c)
// }

// This example demonstrates a simple use of this package.
func Example_simpleUse() {
	s := New(2, 4, 2)
	var a int
	var b float64
	var c int64

	s.Scan("12 2.2 1", &a, &b, &c)
	fmt.Printf("%d %f %d\n", a, b, c)
	// Output: 12 2.200000 1
}

// This example demonstrates that missing fields are filled with its zero-value.
func Example_missingFields() {
	s := New(2, 4, 2)
	var a int
	var b float64
	var c int64

	s.Scan("       1", &a, &b, &c)
	fmt.Printf("%d %f %d\n", a, b, c)
	// Output: 0 0.000000 1
}

// This example shows that it is possible to parse fields with no blanks between them.
func Example_noBlanks() {
	s := New(2, 4, 2)
	var a int
	var b float64
	var c int64

	s.Scan("123.45-1", &a, &b, &c)
	fmt.Printf("%d %f %d\n", a, b, c)
	// Output: 12 3.450000 -1
}
