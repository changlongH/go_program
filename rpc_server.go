package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

type Args struct {
	A int
	B int
}

type Quotient struct {
	Quo int
	Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}


func main() {
	arith := new(Arith)
	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")

	if err != nil {
		fmt.Println("Resolve Tcp Addr", err)
		return;
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("list tcp", err)
		return;
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}

