package main

import (
	"fmt"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
  fmt.Println("New connection")
}

func main() {
  ln, err := net.Listen("tcp", ":8080")

  if err != nil {
    fmt.Println("Failed to bind to port 6379")
    os.Exit(1)

  }

  defer ln.Close()

  for {
    conn, err := ln.Accept()

    if err != nil {
      // handle error
    }
    go handleConnection(conn)
  }


  
}