package main

import (
	"fmt"
)

func main()  {

		var num int = 100
		
		var prt *int = &num

		fmt.Println(&num)

		*prt = 99

		fmt.Println(num)
		
}