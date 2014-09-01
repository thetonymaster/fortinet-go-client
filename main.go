package main

import (
  "antonio-cabreraglz/fortinet-go-client/proxymanager"
)

func main() {
  go proxymanager.StartServer()
}
