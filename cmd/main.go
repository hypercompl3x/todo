package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hypercompl3x/todo/db"
	"github.com/hypercompl3x/todo/handler"
)

func main() {
	router := http.NewServeMux()

	database := db.InitDb()
	defer database.Close()

	fileServer := http.FileServer(http.Dir("./static"))
  router.Handle("GET /static/*", http.StripPrefix("/static/", fileServer))

	todoHandler := handler.TodoHandler{DB: database}

	router.HandleFunc("GET /todos", todoHandler.HandleShowTodo)
	router.HandleFunc("POST /add-todo", todoHandler.HandleAddTodo)
	router.HandleFunc("PUT /update-todo", todoHandler.HandleUpdateTodo)
	router.HandleFunc("GET /show-modal", handler.HandleShowAddModal)
	router.HandleFunc("GET /hide-modal", handler.HandleHideAddModal)

	router.HandleFunc("GET /favicon.ico", http.NotFound)
	router.HandleFunc("GET /*", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/todos", http.StatusSeeOther)
})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server is running on port :8080")
	log.Fatal(server.ListenAndServe())
}