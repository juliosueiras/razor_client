package main

import (
	"fmt"
	"github.com/juliosueiras/razor_client/api"
	"github.com/juliosueiras/razor_client/api/repo"
)

func main() {
	client := razor.New()
	client.SetBaseURL("http://142.55.244.227:8150/api/")

	res, _ := client.Repo.DeleteRepo("test")
	//repo, err := client.Repo.RepoDetails("esxi-60")
	newRepo := &repo.Repo{
		Name: "test",
		URL:  "http://google.ca",
		Task: "tas",
	}

	repoItem, err := client.Repo.CreateRepo(newRepo)

	fmt.Println(err)
	fmt.Println(res)
	fmt.Println(repoItem)

}
