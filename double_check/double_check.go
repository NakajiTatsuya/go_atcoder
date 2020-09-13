package main

import (
	"fmt"
	"strconv"
)

func main() {
	var S string
	fmt.Scan(&S)
	x, err := strconv.Atoi(S)
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(double(x))
	}
}

func double(x int) int {
	return 2 * x
}
