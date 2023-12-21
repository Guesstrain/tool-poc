package main

import (
	"fmt"
	"net"
)

func handleConn(c net.Conn) {
	defer c.Close()
	frameCodec := frame.N
}

func main() {
	l, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	fmt.Println("server start ok(on *.8888)")

	for {
		c, err = l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}

	}
}
