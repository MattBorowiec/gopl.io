// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package popcountShift

import ()

// PopCount returns the population count (number of set bits) of x.
func PopCountShift(x uint64) int {
	var count int
	var i uint64
	for i = 0; i < 64; i++ {
		if (x>>i)&(1<<0) == 1 {
			count += 1
		}
	}
	return count
}

func PopCountShiftRecurse(x uint64) int {
	return popCountShiftRecurse(x, 0)
}
func popCountShiftRecurse(x uint64, count int) int {
	if x == 0 {
		return count
	} else {
		count += int(x & (1 << 0))
		return popCountShiftRecurse(x>>1, count)
	}
}

//!-
