package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// var s, sep string
	for i, arg := range os.Args[1:] {
		//s += sep + arg + strconv.Itoa(i)
		//fmt.Println(i)
		//sep = " "
		//fmt.Println(s)
		fmt.Println(strconv.Itoa(i) + " " + arg)
	}

}
