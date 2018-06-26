package policy

import (
	"github.com/dghubble/sling"
	"github.com/juliosueiras/razor_client/api/error"
	"github.com/juliosueiras/razor_client/api/misc"
)

type PolicyService struct {
	Client *sling.Sling
}

type PolicyItem struct {
	ID   string
	Name string
	Spec string
}

type Policies struct {
	Items []PolicyItem
	Spec  string
	Total int
}

type PolicyRequest struct {
	Name         string                 `json:"name"`
	NodeMetadata map[string]interface{} `json:"node_metadata,omitempty"`
	Broker       string                 `json:"broker"`
	Enabled      bool                   `json:"enabled,omitempty"`
	Before       string                 `json:"before,omitempty"`
	After        string                 `json:"after,omitempty"`
	MaxCount     int                    `json:"max_count,omitempty"`
	RootPassword string                 `json:"root_password"`
	Task         string                 `json:"task,omitempty"`
	Hostname     string                 `json:"hostname"`
	Tags         []string               `json:"tags,omitempty"`
	Repo         string                 `json:"repo"`
}

type NameType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Spec string `json:"spec"`
}

type PolicyNodes struct {
	Count int    `json:"count"`
	ID    string `json:"id"`
	Name  string `json:"name"`
}

type PolicyBroker struct {
	Name string `json:"name"`
}

type Configuration struct {
	HostnamePattern string `json:"hostname_pattern"`
	RootPassword    string `json:"root_password"`
}

type Policy struct {
	Broker        PolicyBroker           `json:"broker"`
	Configuration Configuration          `json:"configuration"`
	NodeMetadata  map[string]interface{} `json:"node_metadata"`
	Enabled       bool                   `json:"enabled"`
	ID            string                 `json:"id"`
	MaxCount      int                    `json:"max_count"`
	Name          string                 `json:"name"`
	Nodes         PolicyNodes            `json:"nodes"`
	Repo          NameType               `json:"repo"`
	Spec          string                 `json:"spec"`
	Tags          []NameType             `json:"tags"`
	Task          NameType               `json:"task"`
}

func (r PolicyService) DeletePolicy(name string) (*misc.ResultMessage, *errorMsg.ErrorMessage) {
	resMsg := new(misc.ResultMessage)
	resError := new(errorMsg.ErrorMessage)

	r.Client.Post("/api/commands/delete-policy").BodyJSON(&struct {
		Name string `json:"name"`
	}{
		Name: name,
	}).Receive(resMsg, resError)

	return resMsg, resError
}

func (r PolicyService) CreatePolicy(policy *PolicyRequest) (*PolicyItem, *errorMsg.ErrorMessage) {
	resErr := new(errorMsg.ErrorMessage)
	policyItem := new(PolicyItem)
	r.Client.Post("/api/commands/create-policy").BodyJSON(policy).Receive(policyItem, resErr)

	return policyItem, resErr
}

func (r PolicyService) UpdatePolicyTask(policyName string, taskName string) (*misc.ResultMessage, *errorMsg.ErrorMessage) {
	resErr := new(errorMsg.ErrorMessage)

	res := new(misc.ResultMessage)
	updatePolicy := &struct {
		Name string `json:"name"`
		Task string `json:"task"`
	}{
		Name: policyName,
		Task: taskName,
	}

	r.Client.Post("/api/commands/update-policy-task").BodyJSON(updatePolicy).Receive(res, resErr)

	return res, resErr
}

func (r PolicyService) ListPolicies() (*Policies, error) {
	policies := new(Policies)
	_, err := r.Client.Get("/api/collections/policies").ReceiveSuccess(policies)

	return policies, err
}

func (r PolicyService) PolicyDetails(id string) (*Policy, *errorMsg.ErrorMessage) {
	policy := new(Policy)
	resErr := new(errorMsg.ErrorMessage)
	r.Client.Get("/api/collections/policies/"+id).Receive(policy, resErr)

	return policy, resErr
}
