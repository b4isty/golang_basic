package main 
 import "fmt"

 func main(){
 	//ARRAY DECLAIRATION
 	// var int_arr [5]int
 	// a := [5]int{1,2,3,4,5}
 	// b := []int{5,6,7,8}

 	var int_arr [5]int
 	fmt.Println(int_arr)

 	var bool_arr [10]bool
 	fmt.Println(bool_arr)

 	var str_arr [15]string
 	fmt.Println(str_arr)

 	int_arr[0]=5
 	int_arr[2]=6
 	fmt.Println(int_arr)
 	fmt.Println(int_arr[0])

 	a := [5]int{1,2,3,4,5}
 	fmt.Println(a)
 	fmt.Println(len(a))
 	fmt.Println(len(bool_arr))
 	b := []int{5,6,7,8}
 	fmt.Println(b)

 	// multi dimentional array 
 	var aa[5][5]int
 	fmt.Println(aa)

 	count := 0

 	for i:=0; i<5; i++{
 		for j:=0; j<5; j++{
 			aa[i][j] = count
 			count = count + 1

 		}
 	}
 	fmt.Println(aa)
 }