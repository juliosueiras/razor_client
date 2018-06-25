package repo

import (
	"fmt"
	"github.com/dghubble/sling"
	"github.com/juliosueiras/razor_client/api/error"
	"github.com/juliosueiras/razor_client/api/misc"
)

type RepoService struct {
	Client *sling.Sling
}

type RepoItem struct {
	ID   string
	Name string
	Spec string
}

type Repos struct {
	Items []RepoItem
	Spec  string
	Total int
}

type Repo struct {
	ID     string `json:"id,omitempty"`
	IsoURL string `json:"iso_url,omitempty"`
	Name   string `json:"name"`
	Spec   string `json:"spec,omitempty"`
	Task   string `json:"task"`
	URL    string `json:"url,omitempty"`
}

func (r RepoService) DeleteRepo(name string) (*misc.ResultMessage, *errorMsg.ErrorMessage) {
	resMsg := new(misc.ResultMessage)
	resError := new(errorMsg.ErrorMessage)

	r.Client.Post("/commands/delete-repo").BodyJSON(&struct {
		Name string `json:"name"`
	}{
		Name: name,
	}).Receive(resMsg, resError)

	return resMsg, resError
}

func (r RepoService) CreateRepo(repo *Repo) (*RepoItem, *errorMsg.ErrorMessage) {
	resErr := new(errorMsg.ErrorMessage)
	repoItem := new(RepoItem)
	req, _ := r.Client.Post("/commands/create-repo").BodyJSON(repo).Request()
	res, err := r.Client.Do(req, repoItem, resErr)
	fmt.Println(res.StatusCode)
	fmt.Println(repo)
	fmt.Println(err)

	return repoItem, resErr
}

func (r RepoService) ListRepos() (*Repos, error) {
	repos := new(Repos)
	_, err := r.Client.Get("/collections/repos").ReceiveSuccess(repos)

	return repos, err
}

func (r RepoService) RepoDetails(id string) (*Repo, *errorMsg.ErrorMessage) {
	repo := new(Repo)
	resErr := new(errorMsg.ErrorMessage)
	r.Client.Get("/collections/repos/"+id).Receive(repo, resErr)

	return repo, resErr
}
