package task

import (
	"github.com/dghubble/sling"
	//"github.com/juliosueiras/razor_client/api/error"
	//"github.com/juliosueiras/razor_client/api/misc"
)

type TaskService struct {
	Client *sling.Sling
}

type TaskItem struct {
	ID   string
	Name string
	Spec string
}

type Tasks struct {
	Items []TaskItem
	Spec  string
}

func (r TaskService) ListTasks() (*Tasks, error) {
	tasks := new(Tasks)
	_, err := r.Client.Get("/api/collections/tasks").ReceiveSuccess(tasks)

	return tasks, err
}

func (r TaskService) CheckIfTaskExist(taskName string) (bool, error) {
	tasks := new(Tasks)
	_, err := r.Client.Get("/api/collections/tasks").ReceiveSuccess(tasks)

	if err != nil {
		return false, err
	}

	for i := range tasks.Items {
		if tasks.Items[i].Name == taskName {
			return true, nil
		}
	}

	return false, nil
}
