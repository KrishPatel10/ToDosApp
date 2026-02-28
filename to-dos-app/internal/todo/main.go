package todo

import (
	"fmt"
)

type Task struct {
	Index       int    `json:"index"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	Deadline    string `json:"deadline"`
	Title       string `json:"title"`
}

type App struct {
	Tasks    map[int]Task
	maxIndex int
}

func NewApp() *App {
	return &App{
		Tasks:    make(map[int]Task),
		maxIndex: 0,
	}
}

func (a *App) AddTask(task Task) bool {
	a.Tasks[task.Index] = task
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
	a.Tasks[task.Index] = task
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

func (a *App) GetAllTasks() []Task {
	tasks := make([]Task, 0, len(a.Tasks))
	for _, task := range a.Tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

func (a *App) DeleteTask(index int) bool {
	if index <= 0 || index >= a.maxIndex {
		return false
	}

	if _, exists := a.Tasks[index]; exists {
		delete(a.Tasks, index)
		return true
	}

	return false
}

func (a *App) MarkTaskAsCompleted(index int) bool {
	if index <= 0 || index > a.maxIndex {
		return false
	}

	if task, exists := a.Tasks[index]; exists {
		task.Completed = true
		a.Tasks[index] = task
		return true
	}
	return false
}

func (a *App) UpdateTask(index int, title, description string) bool {
	if index <= 0 || index > a.maxIndex {
		return false
	}
	if task, exists := a.Tasks[index]; exists {
		task.Title = title
		task.Description = description
		a.Tasks[index] = task
		return true
	}
	return false
}

func (a *App) RemoveTaskByIndex(index int) bool {
	if index <= 0 || index > a.maxIndex {
		return false
	}

	if _, exists := a.Tasks[index]; exists {
		delete(a.Tasks, index)
		return true
	}
	return false
}

func (a *App) findTaskByIndex(index int) *Task {
	if task, exists := a.Tasks[index]; exists {
		return &task
	}
	return nil
}
