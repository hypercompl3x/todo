package handler

import (
	"net/http"

	"github.com/hypercompl3x/todo/services"
	"github.com/hypercompl3x/todo/utils"
	"github.com/hypercompl3x/todo/view"
)

func (h TodoHandler) HandleAddTodo(w http.ResponseWriter, r *http.Request) {
	model, err := services.CreateTodo(h.DB, r.FormValue("title"), r.FormValue("description"))

	if err != nil {
		utils.NewInternalServerError(w, err.Error())
		return
	}

	view.TodoCard(*model).Render(r.Context(), w)
}
