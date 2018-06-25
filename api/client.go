package razor

import (
	"github.com/dghubble/sling"
	"github.com/juliosueiras/razor_client/api/repo"
	"github.com/juliosueiras/razor_client/api/task"
)

type Client struct {
	baseURL    string
	HTTPClient *sling.Sling
	Repo       *repo.RepoService
	Task       *task.TaskService
}

func (c *Client) SetBaseURL(newUrl string) {
	c.HTTPClient.Base(newUrl)
}

func New() *Client {
	c := new(Client)
	c.HTTPClient = sling.New()

	c.Repo = &repo.RepoService{
		Client: c.HTTPClient,
	}

	c.Task = &task.TaskService{
		Client: c.HTTPClient,
	}

	return c
}
