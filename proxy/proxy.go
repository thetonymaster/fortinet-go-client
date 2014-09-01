package proxy

import (
  "fmt"
  "net"
  "time"
  "antonio-cabreraglz/fortinet-go-client/proxymanager"
  "os"
  "syscall"
)

func ListenUDP(addr string) {
  udpAddress, err := net.ResolveUDPAddr("udp4", addr)
  if err != nil {
    panic(err)
  }
  conn, err := net.ListenUDP("udp4", udpAddress)

  if err != nil {
    panic(err)
  }

  defer conn.Close()

  var buf []byte = make([]byte, 1500)
  udpChannel := make(chan []byte)

  for _, address := range proxymanager.GetAddresses() {
    go StartProxyWriter(udpChannel, address)
  }

  for {
    time.Sleep(100 * time.Millisecond)
    n, _, err := conn.ReadFromUDP(buf)

    if err != nil {
      panic(err)
    }

    udpChannel <- buf[0:n]
  }

}

func StartProxyWriter(updChannel <- chan []byte, forwardAddress string) {
  udpWriterAddr, err := net.ResolveUDPAddr("udp4", forwardAddress)

  if err != nil {
    panic(err)
  }

  connWriter, udpErr := net.DialUDP("udp4", nil, udpWriterAddr)
  defer connWriter.Close()

  if udpErr != nil {
    panic(udpErr)
  }

  for msg := range updChannel {
    _, wError := connWriter.Write(msg)
    if wError != nil {
      panic(wError)
    }
  }

}

func StartListener(addr string, sigs chan <- os.Signal){

  udpAddress, err := net.ResolveUDPAddr("udp4", addr)
  if err != nil {
    return
  }

  conn, err := net.ListenUDP("udp4", udpAddress)
  defer conn.Close()

  if err != nil {
    return
  }


  var buf []byte = make([]byte, 1500)

  for {
    time.Sleep(100 * time.Millisecond)
    n, _, err := conn.ReadFromUDP(buf)

    if err != nil {
      panic(err)
    }

    fmt.Println(n)
    fmt.Println(string(buf[0:n]))

    sigs <- syscall.SIGINT
  }
}
