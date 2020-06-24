package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	url := request(conn)
	respond(conn, url)
}

func request(conn net.Conn) string {
	i := 0
	scanner := bufio.NewScanner(conn)
	url := ""
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			m := strings.Fields(ln)
			fmt.Println("****METHOD", m[0])
			fmt.Println("****URL", m[1])
			url = m[1]
		}
		if ln == "" {
			break
		}
		i++
	}
	return url
}

func respond(conn net.Conn, url string) {

	body := `<!DOCTYPE html>
<html>
    <head>
        <!-- head definitions go here -->
    </head>
    <body>
        <strong>URL : %v</strong>
    </body>
</html>`
	urlHtml := fmt.Sprintf(body, url)
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, urlHtml)
}
