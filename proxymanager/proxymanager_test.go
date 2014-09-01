package proxymanager

import (
  "io/ioutil"
  "fmt"
  "os"
)


func ExampleAddIp() {
  CleanUp()

  Execute("a 192.123.123.123\r\n")

  fileBytes, err := ioutil.ReadFile("proxylist.txt")

  if err != nil {
    panic(err)
  }

  fmt.Println(string(fileBytes))

  // Output:
  // 192.123.123.123
}

func ExampleRemoveIp() {
  CleanUp()

  Execute("a 192.123.123.123\r\n")
  Execute("r 192.123.123.123\r\n")

  fileBytes, err := ioutil.ReadFile("proxylist.txt")

  if err != nil {
    panic(err)
  }

  fmt.Println(string(fileBytes))

  // Output:
  // 
}

func ExampleListAll(){
  CleanUp()
  Execute("a 192.168.0.123\r\n")
  Execute("a 192.168.0.124\r\n")

  fmt.Print(Execute("l "))

  _, err := ioutil.ReadFile("proxylist.txt")

  if err == nil {
    panic(err)
  }

  // Output:
  // [192.168.0.123 192.168.0.124]

}

func CleanUp() {
  os.Remove("proxylist.txt")
}

func ExampleCleanCommand() {
  test := "l \r\n"
  fmt.Printf("%s", cleanCommand(test))
  // Output:
  // l
}
