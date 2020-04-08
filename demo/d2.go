package main

import (
	"fmt"
	_ "net/http/pprof"
	"net/url"
)

func main() {
	// url values
	var us = url.Values{}
	us.Set("name", "han")
	us.Add("age", "10")
	us.Add("age", "12")

	fmt.Println(us.Get("name"), us.Get("age"), us.Encode())

	// parse
	u, _ := url.Parse("http://bing.com/search?q=dotnet")
	uu := u.Query()

	fmt.Println(u.Path, u.Host, uu.Get("q"))
}
