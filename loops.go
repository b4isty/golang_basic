package main 

import "fmt"


func main() {
	i:=0
	for i<=10{
		fmt.Println(i)
		i = i + 1
	}

	for j:=0; j<=5; j++{
		if j%2 == 0{
			continue
		}else{
			fmt.Println(j)
		}
	}

	for{
		fmt.Println("keep running")
		break
	}
}