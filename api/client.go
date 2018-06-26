package razor

import (
	"github.com/dghubble/sling"
	"github.com/juliosueiras/razor_client/api/node"
	"github.com/juliosueiras/razor_client/api/policy"
	"github.com/juliosueiras/razor_client/api/repo"
	"github.com/juliosueiras/razor_client/api/tag"
	"github.com/juliosueiras/razor_client/api/task"
)

type Client struct {
	baseURL    string
	HTTPClient *sling.Sling
	Repo       *repo.RepoService
	Task       *task.TaskService
	Node       *node.NodeService
	Tag        *tag.TagService
	Policy     *policy.PolicyService
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

	c.Node = &node.NodeService{
		Client: c.HTTPClient,
	}

	c.Tag = &tag.TagService{
		Client: c.HTTPClient,
	}

	c.Policy = &policy.PolicyService{
		Client: c.HTTPClient,
	}

	return c
}
