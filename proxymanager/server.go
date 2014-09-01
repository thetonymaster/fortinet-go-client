package proxymanager

import (
  "fmt"
  "log"
  "strings"
  "net"
)

func StartServer() {
  ln, err := net.Listen("tcp", ":6000")
  if err != nil {
    log.Fatal(err)
  }

  msgchan := make(chan string)
  resultChannel := make(chan string)

  go executeCommand(msgchan, resultChannel)

  for {
    conn, err := ln.Accept()
    if err != nil {
      log.Println(err)
      continue
    }

    go handleConnection(conn, msgchan, resultChannel)
  }
}

func handleConnection(c net.Conn, msgchan chan <- string, resultChannel <- chan string) {
  fmt.Println("Accepted connection")
  buf := make([]byte, 4096)

  go func(c net.Conn, resultChannel <- chan string) {
    for result := range resultChannel {
      c.Write([]byte(result))
    }
  }(c, resultChannel)


  for {
    n, err := c.Read(buf)
    if err != nil || n == 0 {
      c.Close()
      break
    }

    msgchan <- string(buf[0:n])
  }

  log.Printf("Connection from %v closed.", c.RemoteAddr())
}

func executeCommand(msgchan <-chan string, resultChannel chan <- string){
  for command := range msgchan {
    resultChannel <- "Executing " + command + "\n"
    result := Execute(command)
    resultChannel <- "Result is " + strings.Join(result, ",") + "\n"
  }
}
