package database

import "fmt"

type Task struct {
	ID      string `json:"id,omitempty"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

// Добавление задачи в базу данных
func AddTask(task Task) (int64, error) {
	fmt.Println("444444=")
	res, err := database.Exec(`
		INSERT INTO scheduler (date, title, comment, repeat)
		VALUES (?, ?, ?, ?)
	`, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return 0, err
	}

	var i int64
	i, err = res.LastInsertId()
	if err != nil {

		return 0, err
	}
	fmt.Println("5555=", i)
	return res.LastInsertId()
}
