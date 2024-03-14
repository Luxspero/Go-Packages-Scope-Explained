package main

import (
	"2scope-package/integers"
	"fmt"
)

var five = 5

func main() {
	{
		var integer = 8
		fmt.Println(integer)
	}

	var integer = 1
	fmt.Println(integer)
	fmt.Println(integers.Three)
}
