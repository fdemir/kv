package main

import (
	"fmt"
	"io"
	"net"
	"time"
)
func main() {
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("Error connecting to server", err)
		return
	}

	defer conn.Close()


	conn.Write([]byte("SET a $50\r\n"))

	go func() {
		time.Sleep(1 * time.Second)
		conn.Write([]byte("GET a\r\n"))
	}()

	buff := make([]byte, 1024)

	for {
		n, err := conn.Read(buff)

    if err != nil {
			if err == io.EOF {
					// Connection closed
					break
			} else {
					fmt.Println("Error reading:", err)
					return
			}
	}



		fmt.Println(string(buff[:n]))
	}

}