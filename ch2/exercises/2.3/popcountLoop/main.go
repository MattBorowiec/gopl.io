// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package popcountLoop

import ()

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		//fmt.Printf("IDX: %d, Byte: %8.b\n", i, pc[i/2]+byte(i&1))
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCountLoop(x uint64) int {
	var count int
	var i uint64
	for i = 0; i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}

//!-
