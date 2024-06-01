package handler

import (
	"net/http"

	"github.com/hypercompl3x/todo/view/todo"
)

type TodoHandler struct {}

func (h TodoHandler) HandleShowTodo(w http.ResponseWriter, r *http.Request) {
	err := render(r.Context(), w, todo.Layout())

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}