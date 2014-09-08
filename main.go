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
	flag.Parse()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	if runAsServer {
		serverAddress := ":514"
		go proxymanager.StartServer()
		go proxy.ListenUDP(serverAddress)
		Log("Starting server at " + serverAddress)
	} else {

		go proxy.StartListener(clientAddr)
		Log("Starting client at " + clientAddr)
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
