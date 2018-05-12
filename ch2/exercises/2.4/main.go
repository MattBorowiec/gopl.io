package main

import (
	"fmt"
	"github.com/MattBorowiec/gopl.io/ch2/exercises/2.4/popcountShift"
	"os"
	"strconv"
)

func main() {
	i, _ := strconv.ParseUint(os.Args[1], 10, 64)
	fmt.Printf("PC Shift Recurse: %d\n", popcountShift.PopCountShiftRecurse(i))
	fmt.Printf("PC Shift Loop: %d\n", popcountShift.PopCountShift(i))
}
