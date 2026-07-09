package main

import (
	"fmt"
	"time"
)

type Task struct {
	Id        int
	Title     string
	Done      bool
	CreatedAt time.Time
}

type TaskList struct {
	Tasks  []Task
	NextId int
}

func NewTaskList() *TaskList {
	return &TaskList{
		Tasks:  []Task{},
		NextId: 1,
	}
}

func (taskList *TaskList) Add(title string) (Task, error) {
	var task = Task{
		Id:        taskList.NextId,
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
	}

	taskList.NextId++
	taskList.Tasks = append(taskList.Tasks, task)

	return task, save(taskList)
}

func (taskList *TaskList) Delete(id int) error {
	for i, task := range taskList.Tasks {
		if task.Id == id {
			taskList.Tasks = append(taskList.Tasks[:i], taskList.Tasks[i+1:]...)
			save(taskList)
			return nil
		}
	}

	return fmt.Errorf("task not found")
}

func (taskList *TaskList) Done(id int) error {
	for i := range taskList.Tasks {
		if taskList.Tasks[i].Id == id {
			taskList.Tasks[i].Done = true
			save(taskList)
			return nil
		}
	}

	return fmt.Errorf("task %d not found", id)
}

func (taskList *TaskList) List() {
	for _, task := range taskList.Tasks {
		fmt.Println(task)
	}
}
