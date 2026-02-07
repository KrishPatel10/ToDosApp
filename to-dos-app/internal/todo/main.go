package todo

import (
	"fmt"
)

type Task struct {
	Index       int
	Description string
	Completed   bool
	Deadline    string
	Title       string
}

type App struct {
	Tasks []Task
}

func NewApp() *App {
	return &App{
		Tasks: []Task{},
	}
}

func (a *App) AddTask(task Task) bool {
	a.Tasks = append(a.Tasks, task)
	return true
}

func (a *App) AddTaskByDescription(title string, description string, deadline string) bool {
	task := Task{
		Title:       title,
		Description: description,
		Deadline:    deadline,
		Completed:   false,
		Index:       len(a.Tasks) + 1,
	}
	a.Tasks = append(a.Tasks, task)
	return true
}

func (a *App) ListTasks(showAll, showCompleted bool) {
	fmt.Printf("%-8s | %-20s |%-39s |%-15s | %s \n", "Index", "Title", "Description", "Deadline", "Status")
	for _, task := range a.Tasks {
		var status string

		if task.Completed {
			status = "Done"
		} else {
			status = "Incomplete"
		}

		if showAll {
			fmt.Printf("%-8d | %-20s |%-39s |%-15s | %-3s \n", task.Index, task.Title, task.Description, task.Deadline, status)
		} else if showCompleted && task.Completed {
			fmt.Printf("%-8d | %-20s |%-39s |%-15s | %-3s \n", task.Index, task.Title, task.Description, task.Deadline, status)
		} else if !task.Completed {
			fmt.Printf("%-8d | %-20s |%-39s |%-15s | %-3s \n", task.Index, task.Title, task.Description, task.Deadline, status)
		}
	}
}

func (a *App) MarkTaskAsCompleted(index int) bool {
	if index < 0 || index >= len(a.Tasks) {
		return false
	}

	task := &a.Tasks[index-1]

	task.Completed = true

	return true
}
