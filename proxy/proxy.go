package proxy

import (
	"fmt"
	. "github.com/antonio-cabreraglz/fortinet-go-client/logger"
	"github.com/antonio-cabreraglz/fortinet-go-client/proxymanager"
	"net"
	"time"
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

	for {
		time.Sleep(100 * time.Millisecond)
		n, _, err := conn.ReadFromUDP(buf)

		if err != nil {
			panic(err)
		}


		for _, address := range proxymanager.GetAddresses() {
			go StartProxyWriter(buf[0:n], address)
		}
	}

}

func StartProxyWriter(message []byte, forwardAddress string) {
	Log("Starting proxy writer " + forwardAddress)
	udpWriterAddr, err := net.ResolveUDPAddr("udp4", forwardAddress)

	if err != nil {
		Log(fmt.Sprintf("Forwarding resolve error at %s: %s", forwardAddress, err))
		return
	}

	connWriter, udpErr := net.DialUDP("udp4", nil, udpWriterAddr)
	defer connWriter.Close()

	if udpErr != nil {
		Log(fmt.Sprintf("Forwarding dial error at %s: %s", forwardAddress, udpErr))
		return
	}

	_, wError := connWriter.Write(message)
	if wError != nil {
		Log(fmt.Sprintf("Forwarding write error at %s: %s", forwardAddress, udpErr))
		return
	}

}

// StartListener in parameter addr, if there is an error in the address format
// or in when binding the address we are panicking since we want to know what
// is going on ASAP
func StartListener(addr string) {
	Log("Starting listener on: " + addr)

	udpAddress, resolveErr := net.ResolveUDPAddr("udp4", addr)
	if resolveErr != nil {
		panic(resolveErr)
	}

	conn, listenErr := net.ListenUDP("udp4", udpAddress)
	defer conn.Close()

	if listenErr != nil {
		panic(listenErr)
	}

	var buf []byte = make([]byte, 1500)

	for {
		time.Sleep(100 * time.Millisecond)
		n, _, err := conn.ReadFromUDP(buf)

		if err != nil {
			Log(fmt.Sprintf("Error reading %s: %s", addr, err))
		}

		Log(string(buf[0:n]))
	}
}
