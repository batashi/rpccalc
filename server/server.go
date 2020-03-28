package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Server struct{}

var a, b int64

type nums struct {
	a, b int64
}

// func (this *Server) Negate(i int64, reply *int64) error {
// 	*reply = -i
// 	return nil
// }

func (this *Server) AddA(newA int64, reply *int64) error {
	a = newA
	*reply = 0
	return nil
}

func (this *Server) AddB(newB int64, reply *int64) error {
	b = newB
	*reply = 0
	return nil
}

func (this *Server) Addition(newB int64, reply *int64) error {
	*reply = a + b
	return nil
}

func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	server()
}
