// package main means that the package should be considered a binary executable, instead of a library.
package main

import (
	"fmt"
	// Own package from the same module
	"mymodule/mypackage"
	// External package added from the dependencies in this go.mod file
	"github.com/spf13/cobra"
	// External package from another module
	"golang.org/x/example/hello/reverse"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	cmd := cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, Modules!, from mypackage!")

			mypackage.PrintHello()
		},
	}

	fmt.Println("Calling cmd.Excute()")
	cmd.Execute()

	person1 := Person{
		Name: "Carlos",
		Age:  25,
	}

	fmt.Printf("Hi %s, your name backwards is: %s\n",
		person1.Name,
		reverse.String(person1.Name),
	)
	fmt.Printf("You are %v years old, and your age backward is: %v\n",
		person1.Age,
		reverse.Int(person1.Age),
	)
	// note: a comma is needed after the last argument in a function call
	// if the function call is in a new line.
}
