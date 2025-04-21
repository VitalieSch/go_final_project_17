package database

import (
	"fmt"
)

type Task struct {
	ID      string `json:"id,omitempty"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

// Добавление задачи в базу данных
func AddTask(task Task) (int64, error) {

	res, err := database.Exec(`
		INSERT INTO scheduler (date, title, comment, repeat)
		VALUES (?, ?, ?, ?)
	`, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// Функция получения задачи из БД по ID
func GetTaskByID(id string) (Task, error) {
	var t Task
	var err error

	query := "SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?"
	err = database.QueryRow(query, id).Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat)
	if err != nil {
		return t, err
	}

	return t, nil
}

// Функция обновления задачи для БД
func UpdateTask(task *Task) error {
	var err error

	res, err := database.Exec("UPDATE scheduler SET date = ?, title = ?, comment = ?, repeat = ? WHERE id = ?", task.Date, task.Title, task.Comment, task.Repeat, task.ID)
	if err != nil {
		return err
	}
	// метод RowsAffected() возвращает количество записей,к которым была применена SQL команда
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for updating task`)
	}

	return nil
}

//Функция определения max id в таблице scheduler

func LastId() int {

	var lastId int
	err := database.QueryRow("SELECT max(id) FROM scheduler").Scan(&lastId)

	if err != nil {
		panic(err)
	}
	return lastId
}

// Функция удаления задания из БД по ID
func DeleteTaskById(id string) error {

	var err error

	_, err = database.Exec("DELETE FROM scheduler WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

// Функция обновления даты задания
func UpdateDate(task *Task) error {

	res, err := database.Exec("UPDATE scheduler SET date = ? WHERE id = ?", task.Date, task.ID)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("не корректный id")
	}
	return nil
}
