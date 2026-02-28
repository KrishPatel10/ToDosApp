package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/KrishPatel10/ToDosApp/internal/todo"
)

var app *todo.App // global variable

func main() {
	app = todo.NewApp() // initialize once at startup

	app.AddTaskByDescription("HEHE", "LOLO", time.DateOnly)
	app.AddTaskByDescription("HEHE", "LOLO", time.DateOnly)
	app.AddTaskByDescription("HEHE", "LOLO", time.DateOnly)
	app.AddTaskByDescription("HEHE", "LOLO", time.DateOnly)
	app.AddTaskByDescription("HEHE", "LOLO", time.DateOnly)
	app.AddTaskByDescription("HEHE", "LOLO", time.DateOnly)
	app.AddTaskByDescription("HEHE", "LOLO", time.DateOnly)

	http.HandleFunc("/", listAllTasks)
	http.HandleFunc("PUT /task/{id}", updateTask)
	http.HandleFunc("DELETE /task/{id}", deleteTask)

	http.ListenAndServe(":8080", nil)
}

func listAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := app.GetAllTasks()

	jsonTasks, err := json.Marshal(tasks)

	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonTasks)
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	var task todo.Task
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
	}

	s, err := json.Marshal(task)

	fmt.Print(s)

	if app.UpdateTask(task.Index, task.Title, task.Description) && err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write(s)
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	index, err := strconv.Atoi(strings.TrimSpace(id))

	if err == nil {
		result, _ := json.Marshal(app.DeleteTask(index))
		w.Write(result)
	}
}
