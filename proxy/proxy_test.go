package proxy

import (
  "net"
  "bufio"
  "os"
  "os/signal"
  "syscall"
  "time"
)


func ExampleForwarder() {
  readerServerAddress := ":3333"

  f, err := os.Create("proxylist.txt")
  if err != nil {
    panic(err)
  }
  defer f.Close()

  w := bufio.NewWriter(f)
  _, wrtErr := w.WriteString(readerServerAddress)

  if wrtErr != nil {
    panic(wrtErr)
  }

  w.Flush()

  writerServerAddress := ":3030"

  go ListenUDP(writerServerAddress)
  sigs := make(chan os.Signal, 1)
   signal.Notify(sigs, syscall.SIGINT)
  go StartListener(readerServerAddress, sigs)
  time.Sleep(200 * time.Millisecond)
  //Get the server ready to read packages from the proxy
  // Write to the proxy 
  udpWriterAddr, err := net.ResolveUDPAddr("udp4", writerServerAddress)

  if err != nil {
    panic(err)
  }

  connWriter, udpErr := net.DialUDP("udp4", nil, udpWriterAddr)

  if udpErr != nil {
    panic(udpErr)
  }

  connWriter.Write([]byte("holi"))

  for {
    select {
      case <- sigs: {
        return
      }
    }
  }
  // Output:
  // 4
  // holi
}
