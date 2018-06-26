package tag

import (
	"github.com/dghubble/sling"
	"github.com/juliosueiras/razor_client/api/error"
	"github.com/juliosueiras/razor_client/api/misc"
)

type TagService struct {
	Client *sling.Sling
}

type TagItem struct {
	ID   string
	Name string
	Spec string
}

type Tags struct {
	Items []TagItem
	Spec  string
	Total int
}

type TagRequest struct {
	Name string        `json:"name"`
	Rule []interface{} `json:"rule"`
}

type Count struct {
	Count int    `json:"count"`
	ID    string `json:"id"`
	Name  string `json:"name"`
}

type Tag struct {
	ID       string        `json:"id"`
	Name     string        `json:"name"`
	Policies Count         `json:"policies"`
	Nodes    Count         `json:"nodes"`
	Rule     []interface{} `json:"rule"`
}

func (r TagService) DeleteTag(name string) (*misc.ResultMessage, *errorMsg.ErrorMessage) {
	resMsg := new(misc.ResultMessage)
	resError := new(errorMsg.ErrorMessage)

	r.Client.Post("/api/commands/delete-tag").BodyJSON(&struct {
		Name string `json:"name"`
	}{
		Name: name,
	}).Receive(resMsg, resError)

	return resMsg, resError
}

func (r TagService) CreateTag(tag *TagRequest) (*TagItem, *errorMsg.ErrorMessage) {
	resErr := new(errorMsg.ErrorMessage)
	tagItem := new(TagItem)
	r.Client.Post("/api/commands/create-tag").BodyJSON(tag).Receive(tagItem, resErr)

	return tagItem, resErr
}

func (r TagService) UpdateTagRule(tagName string, rule []interface{}) (*misc.ResultMessage, *errorMsg.ErrorMessage) {
	resErr := new(errorMsg.ErrorMessage)

	res := new(misc.ResultMessage)
	updateTag := &struct {
		Name  string        `json:"name"`
		Rule  []interface{} `json:"rule"`
		Force bool          `json:"force"`
	}{
		Name:  tagName,
		Rule:  rule,
		Force: true,
	}

	r.Client.Post("/api/commands/update-tag-rule").BodyJSON(updateTag).Receive(res, resErr)

	return res, resErr
}

func (r TagService) ListTags() (*Tags, error) {
	tags := new(Tags)
	_, err := r.Client.Get("/api/collections/tags").ReceiveSuccess(tags)

	return tags, err
}

func (r TagService) TagDetails(id string) (*Tag, *errorMsg.ErrorMessage) {
	tag := new(Tag)
	resErr := new(errorMsg.ErrorMessage)
	r.Client.Get("/api/collections/tags/"+id).Receive(tag, resErr)

	return tag, resErr
}
