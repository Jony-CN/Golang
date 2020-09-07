package main

import (
	"fmt"
)

func main()  {

	var n1 float64 = 100
	var n2 float32 = float32(n1)

	fmt.Printf("%T %T",n1,n2)

}