package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Меняем сигнатуры обработчика home, чтобы он определялся как метод
// структуры *application.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w) // Использование помощника notFound()
		//http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		// Поскольку обработчик home теперь является методом структуры application
		// он может получить доступ к логгерам из структуры.
		// Используем их вместо стандартного логгера от Go.
		//app.errorLog.Println(err.Error())
		//http.Error(w, "Internal Server Error", 500)
		app.serverError(w, err) // Использование помощника serverError()
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		// Обновляем код для использования логгера-ошибок
		// из структуры application.
		// app.errorLog.Println(err.Error())
		// http.Error(w, "Internal Serve Error", 500)
		app.serverError(w, err) // Использование помощника serverError()
		return
	}
}

// Меняем сигнатуру обработчика showSnippet, чтобы он был определен как метод
// структуры *application
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		// http.NotFound(w, r)
		app.notFound(w) // Использование помощника notFound()
		return
	}

	fmt.Fprintf(w, "Отображение определенной заметки с ID %d...", id)
}

// Меняем сигнатуру обработчика createSnippet, чтобы он определялся как метод
// структуры *application.
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// http.Error(w, "Метод не дозволен", 405)
		app.clientError(w, http.StatusMethodNotAllowed) // Используем помощник clientError()
		return
	}

	_, err := w.Write([]byte("Создание новой заметки..."))
	if err != nil {
		return
	}
}
