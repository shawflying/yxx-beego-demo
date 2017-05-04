package main

import "fmt"

type Human struct {
}

type Student struct {
	Human //匿名字段
}

type Employee struct {
	Human //匿名字段
}

//在human上面定义了一个method，Student和Employee将继承SayHi方法
func (h *Human) SayHi() {
	fmt.Println("hello")
}

// 重写Human中的SayHi方法
func (s *Student) SayHi(t string) {
	fmt.Println(t)
}

func main() {
	h := Human{}
	h.SayHi();
	e := Employee{}
	e.SayHi()

	s:=Student{}
	s.SayHi("hello world!")

}
