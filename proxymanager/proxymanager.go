package proxymanager

import (
	"io/ioutil"
	"os"
	"strings"
)

func Execute(command string) []string {
	ips := GetAddresses()
	parsedCommand := strings.Split(cleanCommand(command), " ")

	if parsedCommand[0] == "a" {
		ips = append(ips, parsedCommand[1])
		ioutil.WriteFile("proxylist.txt", []byte(strings.Join(ips, ",")), os.ModePerm)
		return nil
	}

	if parsedCommand[0] == "r" {
		for index, value := range ips {
			if value == parsedCommand[1] {
				copy(ips[index:], ips[index+1:])
				ips[len(ips)-1] = ""
				ips = ips[:len(ips)-1]
				break
			}
		}

		ioutil.WriteFile("proxylist.txt", []byte(strings.Join(ips, ",")), os.ModePerm)
		return nil
	}

	if parsedCommand[0] == "l" {
		return ips
	}

	return []string{"Command not understood"}
}

func cleanCommand(command string) string {
	return strings.Trim(strings.Replace(command, "\r\n", " ", -1), " ")
}

func GetAddresses() []string {

	fileContent, err := ioutil.ReadFile("proxylist.txt")
	if err != nil {
		ioutil.WriteFile("proxylist.txt", nil, os.ModePerm)
	}

	var ips []string

	if len(fileContent) == 0 {
		ips = []string{}
	} else {
		stringFileContent := string(fileContent)
		ips = strings.Split(stringFileContent, ",")
	}

	return ips
}
