package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		s   string
		sep string
	)
	for i, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		i += 1
		fmt.Println(i, ":", arg)
	}
	fmt.Println(s)
}
