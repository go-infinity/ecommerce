package main

import (
	"fmt"
)

func main() {
	// string internally is a slice of int - ASCII - value
	var s string = "Azazul" // replace 2 index letter by j

	new := s[:1] + "j" + s[2:]
	// type casting
	fmt.Println(new)
}
