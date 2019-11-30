package main

import (
	"fmt"
	"github.com/gojou/serverpractice/pkg/servers"
)

func main() {
	name := servers.User.get().Name
	fmt.Println("Hello world!")
}
