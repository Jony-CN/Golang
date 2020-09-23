package main

import (
	"fmt"
)

func main()  {

	var a int = 1>>2
	var b int = -1>>2
	var c int = 1<<2
	var d int = -1<<2
// a,b,c,d 的结果是多少

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	

}
