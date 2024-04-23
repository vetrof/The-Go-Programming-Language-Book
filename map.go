package main

import (
	"fmt"
)

type KeyValueCombine struct {
	IntValue 	int
	StringValue string
}


func main() {
	m := make(map[string]KeyValueCombine)
	m["assa"] = KeyValueCombine{123, "jopa"}

	x := m["assa"][0]
	
	fmt.Println(m)
	fmt.Println(x)
}