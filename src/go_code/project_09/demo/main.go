package main

import (
	"fmt"
)

func main()  {
	
	var age int

	fmt.Println("请输入你的年龄！")
	
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("已满18！")
	} else {
		fmt.Println("未满18！")
	}

}