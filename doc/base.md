# 常用响应方法

## this.Ctx.WriteString("ok")

## 返回json数据

```
myjson :=`{"status":"200"}`
	this.Data["json"] =&myjson
	this.ServeJSON()
```

## 默认支持 html 和tpl

## 日志

## 功能
1. 在同一个目录下，package 是一样的
2. 方法函数，参数，大写的表示公开的 小写的表示隐藏的