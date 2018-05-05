// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"strings"
	"testing"
)

func Echo1(str ...string) string {
	var s, sep string
	for i := 0; i < len(str); i++ {
		s += sep + str[i]
		sep = " "
	}
	return s
}

func Echo2(str ...string) string {
	s, sep := "", ""
	for _, arg := range str {
		s += sep + arg
		sep = " "
	}
	return s
}

func Echo3(s ...string) string {
	return strings.Join(s, " ")
}

//!+

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo1("echo1", "arg0", "arg1")
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo2("echo2", "arg0", "arg1")
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo3("echo3", "arg0", "arg1")
	}
}

//!-
