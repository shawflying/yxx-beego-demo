package main

import (
	_ "mohoo-activity-bigwheel/routers"
	_ "github.com/astaxie/beego/session/redis"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
)


func main() {
	logs.SetLogger("console")
	logs.SetLogger(logs.AdapterFile,`{"filename":"logs/project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)

	l := logs.GetLogger()
	l.Println("this is a message of http")
	//an official log.Logger with prefix ORM
	logs.GetLogger("ORM").Println("this is a message of orm")

	logs.Debug("my book is bought in the year of ", 2016)
	logs.Info("this %s cat is %v years old", "yellow", 3)
	logs.Warn("json is a type of kv like", map[string]int{"key": 2016})
	logs.Error(1024, "is a very", "good game")
	logs.Critical("oh,crash")
	beego.Run()
}
