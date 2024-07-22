package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func multipleReturns() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
