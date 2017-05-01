package main

import "fmt"

var a = "w3cshool"
var b string = "b"
var c bool

const mycon string = "abc"

func init() {
	//初始化函数，优先执行
	fmt.Println("init()函数优先执行！")
	title := "我的go小程序"
	desc := "，他是一个简单的小程序"
	fmt.Println(title + desc)
	if title == "我的go小程序" {
		fmt.Println("匹配成功")
	}

	fmt.Println("字体长度")
	fmt.Println(len(title))
	mytitle := &title;
	println(mytitle)
	println(a, b, c) //为什么这个优先执行

	aa, _, cc := 5, 7, "abc"
	println(aa, cc)
}

func main() {
	/**
		这是我的第一个简单的小程序
	 */
	fmt.Println("显示效果")
	fmt.Println("提示时候并不像想象的那么好")
	admin := "显示我们想要的"
	fmt.Println(admin)

	myjson := `{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": ""}`
	fmt.Println(myjson)
}
