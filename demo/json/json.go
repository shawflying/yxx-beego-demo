package main

import (
	"encoding/json"
	"fmt"
)

type Ha struct {
	Iss int `json:"iss"`
	Iat int `json:"iat"`
}
//没有指定结构体，会按照字母排序输出的，可以像下面这样对json指定接收结构体，
func toJSON()  {
	src := `{"iss":148700047,"iat":1487000471}`
	var params Ha
	err := json.Unmarshal([]byte(src), &params)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	p, _ := json.Marshal(params)
	fmt.Printf("%v\n", params)
	fmt.Printf("%s\n", p)
}

func toWxJSON()  {
	src := `{"iss":"magic-ios-1.0.0","iat":1487000471,"jti":"4861ed10","rbd":"dbac2c5f","uri":"/auth/hello","ttu":""}`
	var params interface{}
	err := json.Unmarshal([]byte(src), &params)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	p, _ := json.Marshal(params)
	fmt.Printf("%v, %s\n", params, p)
}

func main()  {
	toWxJSON();
	fmt.Println("-------- 有序")
	toJSON();
}
