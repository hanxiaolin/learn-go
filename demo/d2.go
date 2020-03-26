package main

import (
	"fmt"
	_ "net/http/pprof"
	"net/url"
)

func main() {
	var us = url.Values{}
	us.Set("name", "han")
	us.Add("age", "10")
	us.Add("age", "12")

	fmt.Println(us.Get("name"), us.Get("age"))
}
