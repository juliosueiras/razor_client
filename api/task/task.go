package task

import (
	"github.com/dghubble/sling"
	"github.com/juliosueiras/razor_client/api/error"
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

type TaskRequest struct {
	BootSeq   map[string]string `json:"boot_seq"`
	Name      string            `json:"name"`
	OS        string            `json:"os"`
	Templates string            `json:"templates"`
}

type Task struct {
	BootSeq     map[string]string `json:"boot_seq"`
	Description string            `json:"description"`
	Name        string            `json:"name"`
	OS          struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"os"`
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

func (r TaskService) CreateTask(task *TaskRequest) (*TaskItem, *errorMsg.ErrorMessage) {
	resErr := new(errorMsg.ErrorMessage)
	taskItem := new(TaskItem)
	r.Client.Post("/api/commands/create-task").BodyJSON(task).Receive(taskItem, resErr)

	return taskItem, resErr
}

func (r TaskService) TaskDetails(id string) (*Task, *errorMsg.ErrorMessage) {
	task := new(Task)
	resErr := new(errorMsg.ErrorMessage)
	r.Client.Get("/api/collections/tasks/"+id).Receive(task, resErr)

	return task, resErr
}
