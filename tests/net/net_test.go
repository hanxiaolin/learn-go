package net

import (
	"bufio"
	"fmt"
	"net"
	"testing"
)

func TestA(t *testing.T) {
	//conn, err := net.Dial("tcp", "www.xxx.cn:80")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Fprintf(conn, "GET / HTTP/1.1\r\n\r\n")
	//
	//status, err := bufio.NewReader(conn).ReadString('\n')
	//
	//fmt.Println(status, 33)

	ln, err := net.Listen("tcp", ":8001")
	if err != nil {
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			return
		}

		status, err := bufio.NewReader(conn).ReadString('\n')

		fmt.Println(status, 44)
	}
}
