package main

import (
	"fmt"
	"github.com/juliosueiras/razor_client/api"
)

func main() {
	client := razor.New()

	exist, err := client.Task.CheckIfTaskExist("vmware_esxi2")

	fmt.Println(err)
	fmt.Println(exist)
}
