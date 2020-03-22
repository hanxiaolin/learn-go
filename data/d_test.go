package data

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
	"unsafe"
)

const url = "https://www.baidu.com"

func TestAddxxxx(t *testing.T) {

}

func BenchmarkAddxxxx(b *testing.B) {

}

func TestBulk(t *testing.T) {
	a := "韩晓林2想"
	b := []rune(a)[:10]
	c := []byte(string(b))
	fmt.Println(c)

	var d []byte
	for _, value := range c {
		if value > 0 {
			d = append(d, value)
		}
	}

	fmt.Println(string(a), string(d), 444)
}

func TestXk(t *testing.T) {
	var a = new(struct {
		m map[string]int
		sync.RWMutex
	})
	a.m = map[string]int{}

	var b = struct {
		m map[string]int
		sync.RWMutex
	}{m: map[string]int{}}

	//a.RLock()
	//a.m["mnike"] = 2
	//a.RUnlock()

	a.m["nike"] = 22
	b.m["nike"] = 22
	fmt.Println(a.m, b.m)
}

func TestAdd2(t *testing.T) {
	// int
	var a int
	var b = new(int)
	var aa = new([2]byte)
	var aaa = [2]byte{}
	var bb = make([]byte, 2)
	aa[0] = 1
	aaa[0] = 1
	bb[0] = 1

	fmt.Println(a, *b, *aa, bb, reflect.TypeOf(aa), reflect.TypeOf(bb), reflect.TypeOf(aaa))

	// string
	var c string
	var d = new(string)
	fmt.Println(c, *d, c == *d)

	// map
	var f = map[string]int{}
	f["han"] = 22
	var ff = make(map[string]int)
	ff["han"] = 22
	fmt.Println(f, ff, reflect.TypeOf(ff), reflect.TypeOf(f))

	// slice
	g := make([]int, 2)
	var gg = []int{0}
	gg[0] = 88
	fmt.Println(g[:2], gg)

	var ggg = make([]int, 1)
	copy(ggg, g[0:1])
	ggg[0] = 999

	fmt.Println(ggg, g, "kkk")

	// array
	var e [2]int
	e[0] = 1
	fmt.Println(e)
	var ee = [5]byte{1, 2: 4, 1: 200}
	fmt.Println(e, ee, "xx")

	// struct
	var h struct {
		Name int
	}
	h.Name = 32

	var x = struct {
		Name int
	}{Name: 3232}

	var xx = new(struct{ Name int })
	(*xx).Name = 232
	fmt.Println(h, x, *xx, unsafe.Alignof(x))

}
