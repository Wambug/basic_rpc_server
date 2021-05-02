package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}
type Quotient struct {
	Quo, rem int
}
type Arith int

func (a *Arith) Mulitply(args Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}
func (a *Arith) Divide(args Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("Divide by zero error")
	}
	quo.Quo = args.A / args.B
	quo.rem = args.A % args.B

	return nil

}
func main() {
	arith := new(Arith)
	err := rpc.Register(arith)
	if err != nil {
		log.Fatal("error registering Arith", err)
	}
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":3060")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	log.Printf("serving rpc on port %d", 3060)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("error serving:", err)
	}

}
