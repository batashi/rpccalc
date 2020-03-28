package main

import (
	"fmt"
	"net/rpc"
)

func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	var a, b, result int64
	a = 43
	b = 63
	err = c.Call("Server.AddA", a, &result)
	err = c.Call("Server.AddB", b, &result)

	err = c.Call("Server.Addition", int64(0), &result)
	//fmt.Println(result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server.Addition =", result)
	}
}

func main() {
	client()
}
