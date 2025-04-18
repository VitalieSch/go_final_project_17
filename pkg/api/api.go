package api

import (
	"go1f/pkg/database"
	"net/http"
)

func Init() {

	http.HandleFunc("/api/nextdate", database.NextDateHandler)
	http.HandleFunc("/api/task", TaskHandler)
	http.HandleFunc("/api/tasks", GetTasksHandler)

}
