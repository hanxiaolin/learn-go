package main

import (
	"hanxiaolin/gin-demo/logging"
	"hanxiaolin/gin-demo/models"
	"hanxiaolin/gin-demo/pkg/gredis"
	"hanxiaolin/gin-demo/pkg/setting"
	"hanxiaolin/gin-demo/routers"
	"fmt"
	"github.com/fvbock/endless"
	"log"
	"syscall"
)

type argInt []int

func main() {
	var (
		err error
	)
	//rpc
	//rpcInit()

	//.GET("/proto", func(c *gin.Context) {
	//	person := &pb.Person{
	//		Name:   "yz",
	//		Age:    11,
	//		Emails: []string{"782365461@qq.com", "123456@163.com"},
	//	}
	//
	//	// Marshal
	//	data, err := proto.Marshal(person)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//
	//	fmt.Println(data, 11)
	//
	//	// unMarshal
	//	newObj := new(pb.Person)
	//	err = proto.Unmarshal(data, newObj)
	//	if err != nil {
	//		return
	//	}
	//
	//	c.JSON(200, gin.H{
	//		"xxx": newObj,
	//	})
	//})
	setting.Setup()
	models.Setup()
	logging.Setup()
	err = gredis.Setup()
	if err != nil {
		log.Printf("Server err: %v", err)
	}

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}

//
//func rpcInit() {
//	rpc.Serv()
//}
