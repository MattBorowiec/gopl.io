package main

import (
	"bufio"
	"fmt"
	"github.com/MattBorowiec/gopl.io/ch2/exercises/2.2/units"
	"gopl.io/ch2/tempconv"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Type 'exit' to quit this hot mess")
	if len(os.Args[1:]) > 0 {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(t)
			f := tempconv.Fahrenheit(t)
			c := tempconv.Celsius(t)
			p := units.Pounds(t)
			k := units.Kilograms(t)
			ft := units.Feet(t)
			m := units.Meters(t)
			fmt.Printf("%s = %s\n%s = %s\n%s = %s\n%s = %s\n%s = %s\n%s = %s\n",
				c, tempconv.CToF(c),
				f, tempconv.FToC(f),
				p, units.PToK(p),
				k, units.KToP(k),
				ft, units.FToM(ft),
				m, units.MToF(m))
		}
	} else {
		reader := bufio.NewReader(os.Stdin)
		for {
			input, _ := reader.ReadString('\n')
			input = strings.TrimRight(input, "\n")
			if input == "exit" {
				os.Exit(0)
			}
			in, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("Enter a number you sillyboi")
			}

			f := tempconv.Fahrenheit(in)
			c := tempconv.Celsius(in)
			p := units.Pounds(in)
			k := units.Kilograms(in)
			ft := units.Feet(in)
			m := units.Meters(in)
			fmt.Printf("%s = %s\n%s = %s\n%s = %s\n%s = %s\n%s = %s\n%s = %s\n",
				c, tempconv.CToF(c),
				f, tempconv.FToC(f),
				p, units.PToK(p),
				k, units.KToP(k),
				ft, units.FToM(ft),
				m, units.MToF(m))
		}
	}
}
