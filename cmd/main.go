package main

import (
	"fmt"
	"serverpractice/pkg/entities/scout"
)

func main() {
	scout := new(scout.Scout)
	scout.ID = "mpoling"
	scout.Name = "Mark Poling"

	fmt.Println("Hello" + scout.Name)
}
