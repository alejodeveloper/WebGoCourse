// Problem https://github.com/GoesToEleven/golang-web-dev/tree/master/022_hands-on/02/03_hands-on

package main

import (
	"fmt"
	"bufio"
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
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you said: %s\n", ln)
	}

	io.WriteString(conn, "I see you connected\n")
	fmt.Println("Code got here")

}