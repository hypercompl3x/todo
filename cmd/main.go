package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hypercompl3x/todo/handler"
)

func main() {
	router := http.NewServeMux()

	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	todoHandler := handler.TodoHandler{}

	router.HandleFunc("/todos", todoHandler.HandleShowTodo)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server is running on port :8080")
	log.Fatal(server.ListenAndServe())
}