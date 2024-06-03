package utils

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/hypercompl3x/todo/view"
)

func RenderInLayout(c context.Context, w http.ResponseWriter, component templ.Component) error {
	return view.Layout(component).Render(c, w)
}