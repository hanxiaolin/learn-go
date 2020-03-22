package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// demo1
	//var int1 = 1
	//var int2 int8 = 1
	//var str1 string = fmt.Sprintf("%d", int1)
	//var str2 = fmt.Sprintf("%d", int2)
	var  str3 string = "韩橡皮筋搭配我觉djwojdpwjdjwp得排位的物dhpwdpwjpdjwpdjpw品的价位平均打排位"
	//var int3 uintptr;
	//int3 = 2222

	//var str3 = "韩"
	//var str3 byte = '晓'
	//var str3 int = '\u6653'

	fmt.Println(unsafe.Sizeof(str3))
	//fmt.Println(reflect.TypeOf(str1),
	//	unsafe.Sizeof(int1), unsafe.Sizeof(int2),
	//	unsafe.Sizeof(str1), unsafe.Sizeof(str2), unsafe.Sizeof(str3), &str3,
	//	reflect.TypeOf(int3),
	//)

	//fmt.Printf("%c %U", str3, str3)

	//a := 143.66
	//b := 14.55
	//c := a - b
	//// c = c * 100
	//fmt.Println(c)
	//fmt.Printf("s = %T\n", a)
	//fmt.Printf("s = %T\n", b)
	//fmt.Printf("s = %f %T\n", c, c)

	// demo2
	//str := "你好 world"
	//fmt.Printf("len(str):%d\n", len(str))
	//fmt.Printf("len(rune(str)):%d\n", len([]rune(str)))

	// demo3
	//var a = [2]int8{3, 4}
	//var b = [...]int8{3, 4}
	////var c = new([]int8)
	//var c = make([]int8, 1)
	//
	//c[0] = 1
	////var arr = [...]int{5: 3, 4, 44}
	//fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c), c)

}
