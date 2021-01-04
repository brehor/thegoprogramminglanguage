package main

import (
	"fmt"
	"os"
)

func main() {
	sep := " "
	for index, arg := range os.Args[1:] {
		fmt.Println(fmt.Sprint(index) + sep + arg)
	}
}
