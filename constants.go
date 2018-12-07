package main

import "fmt"

func main() {
	const pi float64 = 3.48768
	fmt.Println(pi)
	const c int = 100
	fmt.Println(c)
	//Can't reassign constant value
	// c = 50

	var d int = 100
	fmt.Println(d)
	d = 5999
	fmt.Println(d)
}



