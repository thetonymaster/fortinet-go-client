package main

import (
	. "github.com/antonio-cabreraglz/fortinet-go-client/logger"

	"flag"
	"fmt"
	"github.com/antonio-cabreraglz/fortinet-go-client/proxy"
	"github.com/antonio-cabreraglz/fortinet-go-client/proxymanager"
	"os"
	"os/signal"
	"syscall"
)

var runAsServer bool
var clientAddr string

func init() {
	flag.BoolVar(&runAsServer, "server", false, "do something")
	flag.StringVar(&clientAddr, "clientAddr", ":3030", "Server Address")
}

func main() {
	Log("asdfakjfhadsjkfhjdksa")
	flag.Parse()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	if runAsServer {
		serverAddress := "255.255.255.255:514"
		go proxymanager.StartServer()
		go proxy.ListenUDP(serverAddress)
		Log("Starting server at " + serverAddress)
	} else {

		serverAddress := ":3030"
		go proxy.StartListener(":3030", nil)
		Log("Starting client at " + serverAddress)
	}

	for {
		select {
		case msg := <-sigs:
			{
				fmt.Println()
				fmt.Println(msg)
				return
			}
		}
	}
}
