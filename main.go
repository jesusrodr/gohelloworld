package main

import (
	"fmt"

	"github.com/jesusrodr/gohelloworld/config"
)

func main() {
	instance := config.Newconfig()
	fmt.Println("hello world! " + instance.Name)

}
