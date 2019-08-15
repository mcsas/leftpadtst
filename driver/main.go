// This is the starting point for the leftpad program
// because of package main

package main

// this is importing the github.com leftpad package
// If import but don't use a package will get a syntax error
// Cannot have cyclic dependencies - pkga imports pkg b and viceversa
// Figure out leftpad-cycle code and test

import (
	"fmt"
	// For the path, notice that it uses forward slashes, not back slashes.
	// \ is an escape character in Go

	//"github.com/jonbodner/go3hours/s4/my-packages/leftpad"
	"github.com/medero/leftpad"
)

func main() {
	fmt.Println(leftpad.Format("hello", 15))
	fmt.Println(leftpad.Format("goodbye", 15))
	fmt.Println(leftpad.Format("¿Cómo está?", 15))
	fmt.Println(leftpad.Format("Internationalization", 15))
	fmt.Println(leftpad.FormatRune("hello", 15, '😀'))
	fmt.Println(leftpad.FormatRune("goodbye", 15, '😀'))
	fmt.Println(leftpad.FormatRune("¿Cómo está?", 15, '😀'))
}
