package taskmodel

import (
	"database/sql"

	"github.com/Dikamayputraa/todo-proa-go/config"
	"github.com/Dikamayputraa/todo-proa-go/entities"
)

type TaskModel struct {
	db *sql.DB
}

func New() *TaskModel {
	db, err := config.DBconn()
	if err != nil {
		panic(err)
	}

	return &TaskModel{db: db}

}

func (m *TaskModel) FindAll(task *[]entities.Task) error {
	rows, err := m.db.Query("select * from tbl_task")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var tasks entities.Task

		rows.Scan(
			&tasks.Id,
			&tasks.Nama,
			&tasks.Des,
			&tasks.Date,
		)

		*task = append(*task, tasks)
	}

	return nil

}

func (m *TaskModel) Create(task *entities.Task) error {
	_, err := m.db.Exec("insert into tbl_task values (?, ?, ?, ?)", nil, task.Nama, task.Des, task.Date)

	if err != nil {
		return err
	}

	return nil
}

func (m *TaskModel) Find(id int64, task *entities.Task) error {
	return m.db.QueryRow("select * from tbl_task where id = ?", id).Scan(

		&task.Id,
		&task.Nama,
		&task.Des,
		&task.Date,
	)
}

func (m *TaskModel) Update(task *entities.Task) error {
	_, err := m.db.Exec("update tbl_task set nama = ?, des = ?, date = ? where id = ?", task.Nama, task.Des, task.Date, task.Id)
	if err != nil {
		return err
	}

	return nil

}

func (m *TaskModel) Delete(id int64) error {
	_, err := m.db.Exec("delete from tbl_task where id = ?", id)
	if err != nil {
		return err
	}

	return nil

}

// func CheckErr(err error) {
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// }
