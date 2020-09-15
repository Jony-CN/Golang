package one

import (
	"fmt"
)

func one()  {
	var name string

	var age int

	var sal float32

	var isPass bool

	fmt.Println("请输入姓名")
	fmt.Scanln(&name)

	fmt.Println("请输入年龄")
	fmt.Scanln(&age)

	fmt.Println("请输入学分")
	fmt.Scanln(&sal)

	fmt.Println("是否通过")
	fmt.Scanln(&isPass)

	fmt.Println(name,age,sal,isPass)
}