package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		io.WriteString(conn, "\nHey, Sup\n")
		fmt.Fprintln(conn, "Sup im a Server")
		fmt.Fprintf(conn, "%v", "Well, See ya")
		conn.Close()
	}
}