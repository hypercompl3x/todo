package handler

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
)

func render(c context.Context, w http.ResponseWriter, component templ.Component) error {
	return component.Render(c, w)
}
