package main

import (
	"faruzan/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	routers.IndexRoutersInit(r)
	r.Run(":100") //默认端口是8000。括号里可以指定端口号，接收字符串类型，比如":4000"。注意":"不能省略！
}
