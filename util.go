package main

import "fmt"

func chk(e error) {
	if e != nil {
		panic(e)
	}
}
func assert(cond bool, format string, v ...any) {
	if !cond {
		panic(fmt.Errorf(format, v...))
	}
}
