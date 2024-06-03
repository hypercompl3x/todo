package handler

import (
	"net/http"
	"strconv"

	"github.com/hypercompl3x/todo/services"
	"github.com/hypercompl3x/todo/utils"
)

func (h TodoHandler) HandleUpdateTodo(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")
	isComplete := r.FormValue("isComplete") == "true"

	id, err := strconv.Atoi(idStr)
	if err != nil {
			utils.NewInternalServerError(w, "Error parsing id: "+err.Error())
			return
	}

	err = services.UpdateTodo(h.DB, id, isComplete)

	if err != nil {
			utils.NewInternalServerError(w, err.Error())
			return
	}
}