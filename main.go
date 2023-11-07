package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

const (
  PONG = "PONG"
  SET = "SET"
  GET = "GET"
  OK = "OK"
)

const (
  INT = 58 // : integer
  STRING = 36 // $ simple string
 )

type Command struct {
  name string
  args []string
}

func parseCommand(command string) Command {
  var name string
  var args []string

  parts := strings.Split(command, " ")

  name = parts[0]
  args = parts[1:]

  
  return Command{
    name,
    args,
  }
}


var data = sync.Map{}

func set(args []string) {
  key := args[0]
  value := args[1][1:]
  valueType := args[1][0]
  
  var newValue string
  
  if valueType == STRING {
    newValue = value
  } else if valueType == INT {
    newValue = value
  }


  data.Store(key, newValue)

}


func handleConnection(conn net.Conn) {
  buffer := make([]byte, 1024)

  for {
    n, err := conn.Read(buffer)

    if err != nil && err.Error() != "EOF" {
     fmt.Println("Error:", err)
     return
    }

    command := string(buffer[:n])

    parsed := parseCommand(strings.Trim(command, "\r\n"))


    switch parsed.name {
      case SET:
        set(parsed.args)
        conn.Write([]byte(OK + "\r\n"))
      case GET:

        // verbose print map
        // fmt.Printf("%#v\n", data)
        v, ok := data.Load(parsed.args[0])

        if !ok {
          conn.Write([]byte("$-1\r\n"))
          break
        }

        conn.Write([]byte("+" + fmt.Sprintf("%s", v)))
      default:
        conn.Write([]byte(PONG + "\r\n"))
    }

  }
}

func main() {
  ln, err := net.Listen("tcp", ":8080")

  if err != nil {
    fmt.Println("Failed to bind to port 8080", err)
    os.Exit(1)
  }

  defer ln.Close()

  for {
    conn, err := ln.Accept()

    if err != nil {
      // EOF = no more connections
      fmt.Println("Error accepting connection", err)
    }

    go handleConnection(conn)
  }
}