package main

import "fmt"

func myFunc() {
	i := 0
Here: //这行的第一个词，以冒号结束作为标签
	println(i)
	i++

	goto Here //跳转到Here去
}
func main() {
	//myFunc();
	/* for配合range可以用于读取slice和map的数据 */
	// 第一个返回值是key，第二个返回值是value
	// 如果是slice，那么key为下标
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
	fmt.Println("测试goto")
}
