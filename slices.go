package main

import "fmt"

func main() {
	// SLICE DECLARATION
	// x := make([]int, 4, 5)
	// a := make([]int, 3)
	// b := []int {2,3,4,5}
	// fmt.Println(x)
	// fmt.Println(a)
	// fmt.Println(b)


	s := make([]int, 3)
	fmt.Println(s)

	// Are similar to array
	s[0] = 1
	s[1] = 2
	s[2] = 3
	// fmt.Println(s)
	// fmt.Println(s[1])
	// fmt.Println(len(s))

	// append funstion is unique to slices

	s = append(s,4)
	// fmt.Println(s)
	s = append(s, 5,6,7)
	fmt.Println(s)
	fmt.Println(s[1:3])
	fmt.Println(s[:3])
	fmt.Println(s[1:])
	fmt.Println(s[:])

	// concise slice defination

	t := []int{10,20,30,40}
	fmt.Println(t)

	x := s
	fmt.Println(x)
	x[0] = 400
	fmt.Println(x)
	fmt.Println(s)

	z := []int{1,2,3,4,5}
	fmt.Println(z)
	y := make([]int, len(s))
	copy(y, z)
	y[0] = 500
	fmt.Println(y)
	fmt.Println(z)

	ss := make([][]int, 3)
	// fmt.Println(ss)
	// [[0], [1,2], [2,3,4]]

	for i := 0; i<3; i++{
		inner_len := i +1
		ss[i] = make([]int, inner_len)
		for j := 0; j<inner_len; j++{
			ss[i][j] = i + j
		}
	}
	fmt.Println(ss)
}