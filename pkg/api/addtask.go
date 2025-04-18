package api

import (
	"encoding/json"
	"fmt"
	"go1f/pkg/database"
	"net/http"
	"strings"
	"time"
)

// Выбор обработчика в зависимости от метода
func TaskHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		AddTaskHandler(w, r)

	}
}

// Обработчик для добавления задачи
func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("11111=")
	var task database.Task
	fmt.Println("222222=", task)
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, `{"error":"Ошибка декодирования JSON"}`, http.StatusBadRequest)
		return
	}
	fmt.Println("333333=", task)
	now := time.Now()

	if task.Date == "" {
		task.Date = now.Format("20060102")
	}

	t, err := time.Parse("20060102", task.Date)
	if err != nil {
		http.Error(w, `{"error":"неверный формат двты"}`, http.StatusBadRequest)
		return
	}

	if t.Format("20060102") == now.Format("20060102") {
		task.Date = now.Format("20060102")
	} else if t.Before(now) && task.Repeat == "" {
		task.Date = now.Format("20060102")
	} else if t.Before(now) {
		nextDate, err := database.NextDate(now, task.Date, task.Repeat)
		if err != nil {
			http.Error(w, `{"error":"неверный формат двты"}`, http.StatusBadRequest)
			return
		}
		task.Date = nextDate
	} else {
		task.Date = t.Format("20060102")
	}

	if task.Title == "" {
		http.Error(w, `{"error":"не указан заголовок задачи"}`, http.StatusBadRequest)
		return
	}

	part := strings.Split(task.Repeat, " ")

	if part[0] != "y" {
		if part[0] != "d" {
			if part[0] != "" {
				http.Error(w, `{"error":"Неверный формат периодичности задачи"}`, http.StatusBadRequest)
				return
			}
		}
	}

	id, err := database.AddTask(task)
	if err != nil {
		http.Error(w, `{"error":"ошибка при добавлении задачи"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"id":"%d"}`, id)

}
