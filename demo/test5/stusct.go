package main

import (
	"fmt"
	"mohoo-activity-bigwheel/demo/test5/function"
)

//可以这个当中一个类

type Books struct {
	title   string
	author  string
	subject string
	book_id int
	function.AddNum
}

func main() {
	var book1 Books
	//var book2 Books
	book1.title = "go 语言"
	book1.author = "yxxit"
	book1.subject = "sub"
	book1.book_id = 10

	fmt.Println(book1.title)
	fmt.Println("程序启动")
}
