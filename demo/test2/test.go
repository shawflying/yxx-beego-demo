package main

import (
	"fmt"
	"unsafe"
)

func init() {
	fmt.Println("初始化")
	a := "abc"
	b := len(a)
	c := unsafe.Sizeof(a)
	println(a, b, c)
}
func main() {
	var a bool = true
	var b bool = false
	if a && b {
		fmt.Println("第一行 - 条件为true\n")
	}

	if a || b {
		fmt.Println("第二行 - 条件为true\n")
	}

	if a == b {
		fmt.Println("打印出来")
	} else {
		fmt.Println("show")
	}

	fmt.Println("显示结果")

	for i := 0; i < 10; i++ {
		for j := 10; j > 0; j-- {
			fmt.Print(i * j)
		}
		fmt.Print("\n")
	}
}
