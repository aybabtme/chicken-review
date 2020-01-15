package views

import (
	"chicken-review/ui"
	"fmt"
	"io"
	"text/template"
)

var DefaultBaseHTMLContext = BaseHTMLContext{
	GlobPattern: "ui/*.gohtml",
	HTML: func(bodyContent interface{}) ui.HTML {
		return ui.HTML{
			Head: ui.Head{
				FavIcoURL:   "",
				Title:       "Chicken Review",
				Author:      "Antoine Grondin",
				Description: "We review chicken restaurants.",
			},
			Body: ui.Body{Content: bodyContent},
		}
	},
}

type BaseHTMLContext struct {
	GlobPattern string
	HTML        func(bodyContent interface{}) ui.HTML
}

func (htmlctx *BaseHTMLContext) RenderUsing(w io.Writer, contentPattern string, bodyContent interface{}) error {
	baseT, err := template.ParseGlob(htmlctx.GlobPattern)
	if err != nil {
		return fmt.Errorf("parsing base html: %v", err)
	}
	contentT, err := template.Must(baseT.Clone()).ParseGlob(contentPattern)
	if err != nil {
		return fmt.Errorf("parsing reviews html: %v", err)
	}

	html := htmlctx.HTML(bodyContent)
	if err := contentT.ExecuteTemplate(w, "html", html); err != nil {
		return fmt.Errorf("executing template: %v", err)
	}
	return nil
}
