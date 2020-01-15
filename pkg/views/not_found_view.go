package views

import (
	"io"
)

type notFoundView struct {
	htmlctx BaseHTMLContext
}

func NewNotFoundView(htmlctx BaseHTMLContext) View {
	return &notFoundView{htmlctx: htmlctx}
}

func (view notFoundView) ContentType() string {
	return "text/html"
}

func (view notFoundView) Render(w io.Writer) error {
	return view.htmlctx.RenderUsing(w, "ui/defaults/404.gohtml", nil)
}
