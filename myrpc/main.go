package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Shape struct {
	A, B int
}

type Basic int

func (b *Basic) Square(s Shape, reply *int) error {
	*reply = s.A * s.A
	return nil
}

func (b *Basic) Rectangle(s Shape, reply *int) error {
	*reply = s.A * s.B
	return nil
}

func main() {
	basic := new(Basic)
	err := rpc.Register(basic)
	if err != nil {
		log.Fatal("erro registering basic", err)
	}
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":3016")
	if err != nil {
		log.Fatal("Listener error", err)
	}
	log.Printf("Serving on port: %d", 3016)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving on port 3016")
	}
}
