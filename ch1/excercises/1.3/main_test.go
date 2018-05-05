// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"testing"
)

func Echo3(s ...string) {
	fmt.Println(strings.Join(os.Args[1:], " ", s))
}

//!+
func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo3("k", "g", 3)
	}
}

//!-
