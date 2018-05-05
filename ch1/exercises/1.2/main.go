// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.

// GOPL Excercise 1.2
package main

import (
	"fmt"
	"os"
)

//!+
func main() {
	for i, v := range os.Args {
		fmt.Println("Idx: ", i, "Value: ", v)
	}
}

//!-
