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
	Tasks    []Task
	maxIndex int
}

func NewApp() *App {
	return &App{
		Tasks:    []Task{},
		maxIndex: 0,
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
		Index:       a.maxIndex + 1,
	}
	a.maxIndex++
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
	if index <= 0 || index > a.maxIndex {
		return false
	}

	task := a.findTaskByIndex(index)
	if task != nil {
		task.Completed = true
	}

	return true
}

func (a *App) UpdateTask(index int, title, description string) bool {
	if index <= 0 || index > a.maxIndex {
		return false
	}
	task := a.findTaskByIndex(index)
	if task != nil {
		task.Title = title
		task.Description = description
		return true
	}
	return false
}

func (a *App) RemoveTaskByIndex(index int) bool {
	if index <= 0 || index > a.maxIndex {
		return false
	}

	taskIndex := a.findTaskIndexInArray(index)

	if taskIndex != -1 {
		a.Tasks = append(a.Tasks[:taskIndex-1], a.Tasks[taskIndex+1:]...)
	}

	return true
}

func (a *App) findTaskIndexInArray(index int) int {
	for ind, task := range a.Tasks {
		if task.Index == index {
			return ind
		}
	}
	return -1
}

func (a *App) findTaskByIndex(index int) *Task {
	for i := range a.Tasks {
		if a.Tasks[i].Index == index {
			return &a.Tasks[i]
		}
	}
	return nil
}
