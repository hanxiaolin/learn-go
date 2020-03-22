package main

import (
	"code.shihuo.cn/gin-demo/data"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		for {
			log.Println(data.Add("https://www.baidu.com"))
		}
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
}
