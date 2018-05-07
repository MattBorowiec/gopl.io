package main

import (
	"fmt"
	"github.com/MattBorowiec/gopl.io/ch2/exercises/2.3/popcountLoop"
	"os"
	"strconv"
)

func main() {
	i, _ := strconv.ParseUint(os.Args[1], 10, 64)
	fmt.Printf("Pop Counts: %d\n", popcountLoop.PopCountLoop(i))
}
