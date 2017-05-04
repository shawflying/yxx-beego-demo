package main

import (
	"github.com/astaxie/beego/httplib"
	"fmt"
	"unsafe"
	"encoding/json"
	"errors"
)

//判断是否活跃
func GetisAlive2(o string) (bool, error) {
	req := httplib.Get("http://172.16.50.138:18082/mh/1.0/queue/isAlive2?openid=" + o)
	var dat map[string]interface{}
	b, _ := req.Bytes()
	fmt.Println(*(*string)(unsafe.Pointer(&b)))
	if err := json.Unmarshal(b, &dat); err != nil {
		return false, errors.New("解析结构失败");
	}
	fmt.Println(dat)
	if dat["status"] == "200" {
		return true, nil;
	} else {
		return false, errors.New("不在活跃");
	}
}

func demo() {
	req := httplib.Get("http://172.16.50.135/wechat/Custom/Ad/QueryAd?location=packHome&type_name=wechat")
	str, err := req.String();
	if err != nil {
		fmt.Println(err)
	}
	b, _ := req.Bytes()
	fmt.Println(*(*string)(unsafe.Pointer(&b)))
	var data map[string]interface{}
	if err := json.Unmarshal(b, &data); err != nil {
		fmt.Println("异常数据",err)
	}
	fmt.Println("当前状态",data["status"])

	fmt.Println(str)
}

func transMap()  {
	
}

func main() {
	openid := "oKXUCj1MOddnp-sCpGi1J1dg3TyM"
	status, err := GetisAlive2(openid)
	demo()
	fmt.Println("http 请求模块", status, err)
}
