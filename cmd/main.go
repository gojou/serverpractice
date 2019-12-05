package main

import (
	"fmt"
	"serverpractice/pkg/servers"
)

func main() {
	name := servers.User.get().Name
	fmt.Println("Hello world!")
}
