package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatalf("encountered error trying to open udp addr: %s", err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatalf("failed to prepare udp connection: %s", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		line, err := reader.ReadString('\n')
		if line != "" {
			_, err = conn.Write([]byte(line))
			if err != nil {
				log.Println(err)
			}
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
		}

	}
}
