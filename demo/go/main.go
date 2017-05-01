package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //使用多核
	a := [...]int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(a)
	c := make(chan bool, 1) //开启
	go func() {
		fmt.Println("Go Go Go!!!")
		c <- true
		close(c) //关闭
	}()
	fmt.Println("项目启动11")
	//<-c
	for v := range c {
		fmt.Println(v)
	}

	t := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go Go(t, i)
	}

	for i := 0; i < 10; i++ {
		<-t
	}

	fmt.Println("项目启动")
}

func Go(c chan bool, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	//if index == 9 {
	//	c <- true
	//}
}
