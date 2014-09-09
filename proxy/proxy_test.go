package proxy

import (
	"bufio"
	"net"
	"io/ioutil"
	"os"
	"time"
	"testing"
	"regexp"
	"strings"
)

func TestForwarder(t *testing.T) {
	cleanUp()

	clientAddress := ":3333"
	writeTestClientAddress(clientAddress)

	serverAddress := ":3030"
	go ListenUDP(serverAddress)

	go StartListener(clientAddress)

	// Give some time to servers to get up
	time.Sleep(100 * time.Millisecond)

	//Get the server ready to read packages from the proxy
	udpWriterAddr, err := net.ResolveUDPAddr("udp4", serverAddress)

	if err != nil {
		panic(err)
	}

	connWriter, udpErr := net.DialUDP("udp4", nil, udpWriterAddr)

	if udpErr != nil {
		panic(udpErr)
	}

	// Write to the proxy
	connWriter.Write([]byte("holi"))

	// Give some time for the message to get to the servers
	time.Sleep(100 * time.Millisecond)

	logBytes, readErr := ioutil.ReadFile("fortinet.log")

	if readErr != nil {
		t.Fatal(readErr)
	}

	logLines := strings.Split(string(logBytes), "\n")
	matched, _ := regexp.MatchString("\\d{4}/\\d{2}/\\d{2} \\d{2}:\\d{2}:\\d{2} holi", logLines[2])

	if !matched {
		t.Fatal(logLines[2])
	}
}

func writeTestClientAddress(clientAddress string) {
	f, err := os.Create("proxylist.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	_, wrtErr := w.WriteString(clientAddress)

	if wrtErr != nil {
		panic(wrtErr)
	}

	w.Flush()
}

func cleanUp() {
	os.Remove("proxylist.txt")
	os.Remove("fortinet.log")
}
