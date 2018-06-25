package repo

import (
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

type RepoRequest struct {
	IsoURL    string `json:"iso_url,omitempty"`
	Name      string `json:"name"`
	Task      string `json:"task"`
	URL       string `json:"url,omitempty"`
	NoContent bool   `json:"no_content,omitempty"`
}

type Repo struct {
	ID     string `json:"id"`
	IsoURL string `json:"iso_url,omitempty"`
	Name   string `json:"name"`
	Task   struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Spec string `json:"spec"`
	} `json:"task"`
	Spec      string `json:"spec"`
	URL       string `json:"url,omitempty"`
	NoContent bool   `json:"no_content,omitempty"`
}

//"boot_seq": {
//"name": {
//"os": {
//"templates": {

func (r RepoService) DeleteRepo(name string) (*misc.ResultMessage, *errorMsg.ErrorMessage) {
	resMsg := new(misc.ResultMessage)
	resError := new(errorMsg.ErrorMessage)

	r.Client.Post("/api/commands/delete-repo").BodyJSON(&struct {
		Name string `json:"name"`
	}{
		Name: name,
	}).Receive(resMsg, resError)

	return resMsg, resError
}

func (r RepoService) CreateRepo(repo *RepoRequest) (*RepoItem, *errorMsg.ErrorMessage) {
	resErr := new(errorMsg.ErrorMessage)
	repoItem := new(RepoItem)
	r.Client.Post("/api/commands/create-repo").BodyJSON(repo).Receive(repoItem, resErr)

	return repoItem, resErr
}

func (r RepoService) UpdateRepoTask(repoName string, taskName string) (*misc.ResultMessage, *errorMsg.ErrorMessage) {
	resErr := new(errorMsg.ErrorMessage)

	res := new(misc.ResultMessage)
	updateRepo := &struct {
		Repo string `json:"repo"`
		Task string `json:"task"`
	}{
		Repo: repoName,
		Task: taskName,
	}

	r.Client.Post("/api/commands/update-repo-task").BodyJSON(updateRepo).Receive(res, resErr)

	return res, resErr
}

func (r RepoService) ListRepos() (*Repos, error) {
	repos := new(Repos)
	_, err := r.Client.Get("/api/collections/repos").ReceiveSuccess(repos)

	return repos, err
}

func (r RepoService) RepoDetails(id string) (*Repo, *errorMsg.ErrorMessage) {
	repo := new(Repo)
	resErr := new(errorMsg.ErrorMessage)
	r.Client.Get("/api/collections/repos/"+id).Receive(repo, resErr)

	return repo, resErr
}
