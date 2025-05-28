package main

import (
	"fmt"
	"httpfromtcp/internal/request"
	"log"
	"net"
)

const port = ":42069"

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("encountered error trying to listen %s", err)
		return
	}
	defer listener.Close()
	fmt.Println("Listening for TCP traffic on", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Accepted connnection from", conn.RemoteAddr())
		req, err := request.RequestFromReader(conn)
		if err != nil {
			log.Fatalf("failed retrieving request from reader: %s", err)
		}
		fmt.Println("Request line:")
		fmt.Println("- Method:", req.RequestLine.Method)
		fmt.Println("- Target:", req.RequestLine.RequestTarget)
		fmt.Println("- Version:", req.RequestLine.HttpVersion)

		fmt.Println("Headers:")
		for key, value := range req.Headers {
			fmt.Printf("- %s: %s\n", key, value)
		}

		fmt.Println("Body:")
		fmt.Print(string(req.Body), "\n")
		fmt.Println("connection closed")
	}
}
