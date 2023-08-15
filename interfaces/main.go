package interfaces

import "fmt"

func Main() {
	fmt.Println("<------------ Interfaces ----------->")
	r := Rect{width: 10, height: 10}
	// s := Square{side: 10}
	c := Circle{radius: 10}
	// TotalArea(r, c, s) => Error Square isn't a Shape !
	// Implicit interfaces inference...
	fmt.Println("Total Area: ", TotalArea(r, c))
	fmt.Println("Total Perimeter: ", TotalPerimeter(r, c))

	// Empty interface
	Double(2)        // doubling an int
	Double(2.2)      // doubling an float
	Double("Golang") // doubling an string
	fmt.Println("<------------ Interfaces ----------->")
}
