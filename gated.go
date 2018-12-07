package main

import (
	"fmt"
	"net"
	//"strings"
)

func connHandler(conn net.Conn) {
	buf := make([]byte, 4096)
	for {
		read_size, err := conn.Read(buf)
		if err != nil || read_size <= 0 {
			//fmt.Println(read_size, err)
			conn.Close()
			break
		}

		fmt.Println(string(buf[0:read_size]))
		w_size, err := conn.Write(buf[0:read_size])
		if err != nil || w_size <= 0 {
			fmt.Println(w_size, err)
		}
	}

	//fmt.Println("Connection from %v closed.", inputs[0])
}

func main() {
	server, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("Fail to start server.", err)
		return
	}

	fmt.Println("Server started ...")

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Fail to connect.", err)
			break
		}

		go connHandler(conn)
	}

	fmt.Println("close")
}
