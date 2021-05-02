package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}
type Quotient struct {
	Quo, rem int
}

func main() {
	var reply int
	var args []Args
	a := Args{23345, 24893}

	client, err := rpc.DialHTTP("tcp", "localhost:3060")
	if err != nil {
		log.Fatal("Connection error:", err)
	}
	err = client.Call("Arith.Mulitply", a, &reply)
	if err != nil {
		log.Fatal("Multiplication error", err)
	}
	fmt.Println("Arith", args, reply)
	//Asynchrounous call
	quotient := new(Quotient)
	err = client.Call("Arith.Divide", a, quotient)
	if err != nil {
		log.Fatal("Division error")
	}
	fmt.Println("Arith", args, quotient)
}
