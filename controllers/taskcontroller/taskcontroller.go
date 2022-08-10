package taskcontroller

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/Dikamayputraa/todo-proa-go/entities"
	"github.com/Dikamayputraa/todo-proa-go/models/taskmodel"
)

var taskModel = taskmodel.New()

func Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"data": template.HTML(GetData()),
	}

	tmp, _ := template.ParseFiles("views/task/index.html")
	tmp.Execute(w, data)
}

func GetData() string {
	buffer := &bytes.Buffer{}

	tmpl, _ := template.New("data.html").Funcs(template.FuncMap{
		"increment": func(a, b int) int {
			return a + b
		},
	}).ParseFiles("views/task/data.html")

	var task []entities.Task

	err := taskModel.FindAll(&task)
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"task": task,
	}

	tmpl.ExecuteTemplate(buffer, "data.html", data)

	return buffer.String()
}

func GetForm(w http.ResponseWriter, r *http.Request) {

	queryString := r.URL.Query()
	id, err := strconv.ParseInt(queryString.Get("id"), 10, 64)

	var data map[string]interface{}

	if err != nil {
		data = map[string]interface{}{
			"title": "Tambah Data Mahasiswa",
		}
	} else {
		var task entities.Task

		err := taskModel.Find(id, &task)

		if err != nil {
			panic(err)
		}

		data = map[string]interface{}{
			"title": "Edit Data Mahasiswa",
			"task":  task,
		}
	}

	tmp, _ := template.ParseFiles("views/task/form.html")
	tmp.Execute(w, data)
}

func Store(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		var task entities.Task

		task.Nama = r.Form.Get("nama")
		task.Des = r.Form.Get("des")
		task.Date = r.Form.Get("date")

		id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)

		var data map[string]interface{}
		if err != nil {
			//insert data
			err := taskModel.Create(&task)
			if err != nil {
				ResponseError(w, http.StatusInternalServerError, err.Error())
				return
			}

			data = map[string]interface{}{
				"message": "successfully added task",
				"data":    template.HTML(GetData()),
			}
		} else {
			//update data
			task.Id = int(id)

			err := taskModel.Update(&task)
			if err != nil {
				ResponseError(w, http.StatusInternalServerError, err.Error())
				return
			}

			data = map[string]interface{}{
				"message": "task changed successfully",
				"data":    template.HTML(GetData()),
			}

		}

		ResponseJson(w, http.StatusOK, data)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)
	if err != nil {
		panic(err)
	}

	err = taskModel.Delete(id)
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"message": "task deleted successfully",
		"data":    template.HTML(GetData()),
	}

	ResponseJson(w, http.StatusOK, data)
}

func ResponseError(w http.ResponseWriter, code int, message string) {
	ResponseJson(w, code, map[string]string{"error": message})
}

func ResponseJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
