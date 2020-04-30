package main

import (
	"net/http"
)

func main() {
	/**
	  第 1 行，标记当前文件为 main 包，main 包也是 Go 程序的入口包。
	  第 3~5 行，导入 net/http 包，这个包的作用是 HTTP 的基础封装和访问。
	  第 7 行，程序执行的入口函数 main()。
	  第 8 行，使用 http.FileServer 文件服务器将当前目录作为根目录（/目录）的处理器，访问根目录，就会进入当前目录。
	  第 9 行，默认的 HTTP 服务侦听在本机 8080 端口。
	*/
	http.Handle("/",http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8088",nil)
}
