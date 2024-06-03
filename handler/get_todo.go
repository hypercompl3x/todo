package handler

import (
	"database/sql"
	"net/http"

	"github.com/hypercompl3x/todo/services"
	"github.com/hypercompl3x/todo/utils"
	"github.com/hypercompl3x/todo/view"
)

type TodoHandler struct {
	DB *sql.DB
}

func (h TodoHandler) HandleShowTodo(w http.ResponseWriter, r *http.Request) {
	todos, err := services.GetTodos(h.DB)

	if err != nil {
		utils.NewInternalServerError(w, err.Error())
		return
	}

	err = utils.RenderInLayout(r.Context(), w, view.Todo(todos))

	if err != nil {
		utils.NewInternalServerError(w, err.Error())
		return
	}
}

func HandleShowAddModal(w http.ResponseWriter, r *http.Request) {
	view.AddModal().Render(r.Context(), w)
}

func HandleHideAddModal(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(nil))
}