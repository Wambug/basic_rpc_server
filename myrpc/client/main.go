package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Shape struct {
	A, B int
}

func main() {
	var reply int
	var s []Shape
	a := Shape{46, 28}

	client, err := rpc.DialHTTP("tcp", "localhost:3016")
	if err != nil {
		log.Fatal("Connection error", err)
	}
	err = client.Call("Basic.Square", a, &reply)
	if err != nil {
		log.Fatal("Error in finding area of square", err)
	}
	fmt.Println(s, reply)
	err = client.Call("Basic.Rectangle", a, &reply)
	if err != nil {
		log.Fatal("Error in finding area of rectangle", err)
	}
	fmt.Println(s, reply)
}
