package main

import (
  "fmt"
  "servers"
)

func main() {
  name:=servers.User.get().Name
  fmt.Println("Hello world!")
}
