package main

import (
	"fmt"
	"strings"
	"strconv"
	"encoding/xml"
	"os"
)

//消除字符串

//合并，前缀后缀检测，包含，链接，消除空格，分割，去索引
//判断是否包含
func ckContains() {
	s := "hello world"
	t := "hello"
	fmt.Println("是否包含", strings.Contains(s, t), strings.Contains(s, "?"))
	fmt.Println("查看索引位置", strings.Index(s, "0"), strings.Index(s, "e")) //-1 表示未查到
	ss := "1#2#234"
	splitedStr := strings.Split(ss, "#")
	fmt.Println("打印分割功能", splitedStr)

	//合并字符串
	fmt.Println("合并字符串", strings.Join(splitedStr, "@"))

}

func chStrconv() {
	s := "20"
	n := 10
	//字符串转换为数字
	fmt.Println("数字转换为字符串：" + strconv.Itoa(n))
	v, e := strconv.Atoi(s)
	fmt.Println("字符串转换为数字", v, e)

	//解析 有可能返回错误，或者转译错误，所以会有一个err选项
	fmt.Println(strconv.ParseBool("false"))
	fmt.Println(strconv.ParseFloat("3.14", 64))

	//格式化
	fmt.Println("格式化bool", strconv.FormatBool(true))
	fmt.Println("转换数字格式10位转换为8进制", strconv.FormatInt(123, 8))
}

//人物档案
type person struct {
	Name string `xml:"name,attr"`
	Age  int `xml:"age"`
}

func ckXml() {
	p := person{Name: "yxx", Age: 18}
	var data []byte
	var err error
	//if data, err = xml.Marshal(p); err != nil {
	if data, err = xml.MarshalIndent(p, "", " "); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("打印格式：", string(data))

	p2 := new(person)
	if err = xml.Unmarshal(data, p2); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p2) //以指针的方式显示
}

func ckOs()  {
	fmt.Println(os.Args)
}

func main() {
	ckContains()
	fmt.Println("-----------------------------------")
	chStrconv()
	fmt.Println("----------------------------------------")
	ckXml()
	fmt.Println("----------------------------------------")
	ckOs()
	fmt.Println("string 常用函数")
}
