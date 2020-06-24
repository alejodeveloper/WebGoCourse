package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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
	io.WriteString(conn, "In Mem DATABASE\n")
	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)

		switch fs[0] {
		case "GET":
			dictKey := fs[1]
			v := data[dictKey]
			fmt.Fprintf(conn, "%s\n", v)

		case "SET":
			if len(fs) != 3 {
				fmt.Fprintln(conn, "EXPECTED VALUE")
			}
			dictKey := fs[1]
			v := fs[2]
			data[dictKey] = v

		case "DEL":

			dictKey := fs[1]
			delete(data, dictKey)
		default:
			fmt.Fprintln(conn, "Invalid command")
		}
	}
}
