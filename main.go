package main

import (
  "os"
  "os/signal"
  "github.com/antonio-cabreraglz/fortinet-go-client/proxymanager"
  "fmt"
  "syscall"
)

func main(){
   sigs := make(chan os.Signal, 1)

   signal.Notify(sigs, syscall.SIGINT)

   go proxymanager.StartServer()

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


