package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func createAcceptRequest(ln net.Listener, ch chan error) net.Conn {
	conn, err := ln.Accept()
	if err != nil {
		ch<-err
	}
	return conn
}


func handleMessageRequest(conn net.Conn) {
	m,_:=bufio.NewReader(conn).ReadString('\n')

		fmt.Print("Message Received:", string(m))

		r := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")

		text, _ := r.ReadString('\n')

		conn.Write([]byte(text + "\n"))
}
func main() {
	ch := make(chan error,1)
	fmt.Println("Launching Server")

	ln, _ := net.Listen("tcp", ":8081")

	conn := createAcceptRequest(ln,ch);

	for {
		handleMessageRequest(conn)
	}
}