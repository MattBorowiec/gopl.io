// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	tmpcnv "github.com/MattBorowiec/gopl.io/ch2/exercises/2.1/tempconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tmpcnv.Fahrenheit(t)
		c := tmpcnv.Celsius(t)
		k := tmpcnv.Kelvin(t)
		fmt.Printf("%s = %s, %s = %s, %s = %s, %s = %s\n",
			f, tmpcnv.FToC(f), c, tmpcnv.CToF(c), k, tmpcnv.KToC(k), f, tmpcnv.CToK(tmpcnv.FToC(f)))
	}
}

//!-
