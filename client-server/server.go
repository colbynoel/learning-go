package main

import (
  "fmt"
  "net"
)

func main(){
  ln, err := net.Listen("tcp", ":8080")
  if err != nil {
    fmt.Println(err)
    return
  }

  // Accept incoming connections and handle them

  for {
    conn, err := ln.Accept()
    if err != nil {
      fmt.Println(err)
      continue
    }

    go handleConnection(conn)
  }
}

func handleConnection(conn net.Conn) {
  // Close the connection when we're done
  defer conn.Close()

  buf := make([]byte, 1024)
  _, err := conn.Read(buf)
  if err != nil{
    fmt.Println(err)
    return
  }

  fmt.Printf("Received : %s", buf)
}
