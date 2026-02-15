package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/KrishPatel10/ToDosApp/internal/todo"
)

func main() {
	app := todo.NewApp()

	runToDo(app)
}

func runToDo(app *todo.App) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("What you want to do?\n1.Log new Task\n2.Mark as Complete\n3.List All Tasks\n4.Edit Tasks\n5.Exit\n")

		choice, err := reader.ReadString('\n')

		if err == nil {
			switch strings.TrimSpace(choice) {
			case "1":
				{
					fmt.Print("Task Name: ")
					taskName, err := reader.ReadString('\n')
					taskName = strings.TrimSpace(taskName)
					if err == nil {
						fmt.Print("Description: ")
						description, err := reader.ReadString('\n')
						description = strings.TrimSpace(description)
						if err == nil {
							app.AddTaskByDescription(taskName, description, time.DateOnly)
						}
					}
				}
			case "2":
				{
					fmt.Print("Which task you want to mark as completed?")
					taskIndex, err := reader.ReadString('\n')
					index, err := strconv.Atoi(strings.TrimSpace(taskIndex))
					if err == nil {
						app.MarkTaskAsCompleted(index)
					}
				}
			case "3":
				{
					app.ListTasks(true, false)
				}
			case "4":
				{
					fmt.Print("Task Index: ")
					taskInd, err := reader.ReadString('\n')
					taskIndex, err := strconv.Atoi(strings.TrimSpace(taskInd))
					fmt.Println("Are", err, taskInd, taskIndex)
					if err == nil {
						fmt.Print("Task Name: ")
						taskName, err := reader.ReadString('\n')
						taskName = strings.TrimSpace(taskName)
						if err == nil {
							fmt.Print("Description: ")
							description, err := reader.ReadString('\n')
							description = strings.TrimSpace(description)
							if err == nil {
								app.UpdateTask(taskIndex, taskName, description)
							}
						}
					}
				}
			case "5":
				return
			}
		}
	}
}
