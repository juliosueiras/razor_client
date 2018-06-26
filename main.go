package main

import (
	"fmt"
	"github.com/juliosueiras/razor_client/api"
)

func main() {
	client := razor.New()
	client.SetBaseURL("http://142.55.244.227:8150/")

	policy, _ := client.Policy.PolicyDetails("test")
	fmt.Println(policy.NodeMetadata)
}
