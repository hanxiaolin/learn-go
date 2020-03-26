package tests

import (
	"encoding/json"
	"fmt"
	"testing"
)

// json使用案例
func Test_json(t *testing.T) {
	// map
	var map1 = map[string]string{
		"name": "han",
	}
	// encode
	byte1, err := json.Marshal(map1)
	if err != nil {
		return
	}
	fmt.Println(byte1, string(byte1), "map")

	// struct
	type person struct {
		Name string `json:"name"`
	}
	p1 := person{Name: "han"}

	// encode
	byte2, err := json.Marshal(p1)
	if err != nil {
		return
	}

	// decode
	p2 := &person{}
	err = json.Unmarshal(byte2, p2)
	if err != nil {
		return
	}
	
	fmt.Println(p1, p2, string(byte2), "struct")
}

// string to int, int to string
func Test_int_string(t *testing.T) {

}
