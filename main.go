package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Yandex-Practicum/final-project-encoding-go/encoding"
	"github.com/Yandex-Practicum/final-project-encoding-go/utils"
	"github.com/go-chi/chi/v5"
)

func Encode(data encoding.MyEncoder) error {
	return data.Encoding()
}

type Task struct {
	ID          string   `json:"id"`          // ID задачи
	Description string   `json:"description"` // Заголовок
	Note        string   `json:"note"`        // Описание задачи
	Application []string `json:"application"` // Приложения, которыми будете пользоваться
}

var tasks = map[string]Task{
	"1": {
		ID:          "1",
		Description: "Сделать финальное задание темы REST API",
		Note:        "Если сегодня сделаю, то завтра будет свободный день. Ура!",
		Application: []string{
			"VS Code",
			"Terminal",
			"git",
		},
	},

	"2": {
		ID:          "2",
		Description: "Протестировать финальное задание с помощью Postmen",
		Note:        "Лучше это делать в процессе разработки, каждый раз, когда запускаешь сервер и проверяешь хендлер",
		Application: []string{
			"Vs Code",
			"Terminal",
			"git",
			"Postman",
		},
	},
}

func getTask(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func getIdTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	task, ok := tasks[id]
	if !ok {
		http.Error(w, "Задача не найдена", http.StatusNoContent)
		return
	}

	resp, err := json.Marshal(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func postTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var task Task
		var buf bytes.Buffer
		_, err := buf.ReadFrom(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err = json.Unmarshal(buf.Bytes(), &task); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		tasks[task.ID] = task
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}

	return
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	task, ok := tasks[id]
	if !ok {
		http.Error(w, "Задача не найдена", http.StatusBadRequest)
		return
	}
	w.Header().Del(task.ID)
	w.WriteHeader(http.StatusOK)
}

func main() {
	utils.CreateJSONFile()
	utils.CreateYAMLFile()

	jsonData := encoding.JSONData{FileInput: "jsonInput.json", FileOutput: "yamlOutput.yml"}
	err := Encode(&jsonData)
	if err != nil {
		fmt.Printf("ошибка при перекодировании данных из JSON в YAML: %s", err.Error())
	}

	yamlData := encoding.YAMLData{FileInput: "yamlInput.yml", FileOutput: "jsonOutput.json"}
	err = Encode(&yamlData)
	if err != nil {
		fmt.Printf("ошибка при перекодировании данных из YAML в JSON: %s", err.Error())
	}

	r := chi.NewRouter()

	r.Get("/tasks", getTask)            // получение всех задач
	r.Get("/tasks/{id}", getIdTask)     // порлучение задачи по id
	r.Post("/tasks", postTask)          //  отправка на сервер запроса
	r.Delete("/tasks/{id}", deleteTask) // удаление

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}

}
