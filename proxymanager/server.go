package proxymanager

import (
	"fmt"
	"net"
	"strings"

	. "github.com/antonio-cabreraglz/fortinet-go-client/logger"
)

func StartServer() {
	ln, err := net.Listen("tcp", ":6000")
	if err != nil {
		Log(err.Error())
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			Log(err.Error())
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	Log("Accepted connection")
	buf := make([]byte, 4096)

	for {
		n, err := c.Read(buf)
		if err != nil || n == 0 {
			c.Close()
			break
		}

		command := string(buf[0:n])
		c.Write([]byte("Executing " + command + "\n"))
		result := Execute(command)
		c.Write([]byte(strings.Join(result, ", ") + "\n"))
	}

	Log(fmt.Sprintf("Connection from %v closed.", c.RemoteAddr()))
}
