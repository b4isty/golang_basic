package main 

import "fmt"

func main() {
	i := 3
	switch i {
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")	
	}

	j := 5
	switch j {
		case 1,2:
			fmt.Println("one or two")
		case 3:
			fmt.Println("three")
		default:
			fmt.Println("NOt one, two or three")
	}

	switch {
		case j==5:
			fmt.Println(j,"is equal to 5")
		case j<5:
			fmt.Println(j, "is less than 5")
		case j>5:
			fmt.Println(j, "is grater than 5")
	}
}
