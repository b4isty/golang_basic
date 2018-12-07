package main 

import "fmt"

func main() {
	i:=7

	if i%2 == 0{
		fmt.Println(i, "is even")
	}else{
		fmt.Println(i, "is odd")
	}

	k:=10
	if k==10{
		fmt.Println(k, "is 10")
	}

	j:=60
	if j<50{
		fmt.Println(j,"is less than 50")
	}else if j>50{
		fmt.Println(j, "is greater than 50")
	}else{
		fmt.Println(j, "is equal to 50")
	}
}
