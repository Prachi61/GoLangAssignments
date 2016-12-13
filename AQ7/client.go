package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")

	for {
		r := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")

		text, _ := r.ReadString('\n')

		fmt.Fprintf(conn, text + "\n")

		m,_ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: "+m)
	}
}