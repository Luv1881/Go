package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":1222")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return
		}
		do(conn) //add go here to make it concurrent
	}
}

func do(conn net.Conn) {
	buf := make([]byte, 1024)
	conn.Read(buf)
	time.Sleep(2 * time.Second)
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World!\r\n"))
	conn.Close()
}
