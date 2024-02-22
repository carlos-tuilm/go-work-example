package mypackage

import "fmt"

// the capital letter in the function name means that it is exported,
// and can be used by other packages.
func PrintHello() {
	fmt.Println("Hello, Modules!, from mypackage!")
}
