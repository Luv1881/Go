package main

import (
	"fmt"
	"net"
	"os"
)

const port = "1222"

func do(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Something is wrong")
			break
		}
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World!\r\n"))
	}
}

func main() {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Something went wrong")
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Printf("Listening on port %s\n", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Cant Accept")
			continue
		}
		go do(conn)
	}
}
