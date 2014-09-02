package main

import (
  "os"
  "os/signal"
  "github.com/antonio-cabreraglz/fortinet-go-client/proxymanager"
  "fmt"
)

var runAsServer bool
var clientAddr string

func init() {
  flag.BoolVar(&runAsServer, "server", false, "do something")
  flag.StringVar(&clientAddr, "clientAddr", ":3030", "Server Address")
}

func main(){
  flag.Parse()

  sigs := make(chan os.Signal, 1)
  signal.Notify(sigs, syscall.SIGINT)

  if runAsServer {
    go  proxy.StartListener(":3030", sigs)
  } else {
    go proxymanager.StartServer()
    go proxy.ListenUDP("255.255.255.255:514")
  }

  for {
    select {
      case  msg:= <- sigs: {
        fmt.Println()
        fmt.Println(msg)
        return
      }
    }
  }
}


