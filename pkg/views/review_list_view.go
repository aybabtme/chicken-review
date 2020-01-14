package views

import (
	"bytes"
	"chicken-review/pkg/models"
	"fmt"
	"io"
)

type ReviewListView struct {
	Reviews []models.Review
}

func (view ReviewListView) ContentType() string {
	return "text/html"
}

func (view ReviewListView) Render(w io.Writer) error {
	html := bytes.NewBuffer(nil)
	html.WriteString("<html>")
	html.WriteString("	<body>")
	html.WriteString("		<h1>List of past reviews!</h1>")
	html.WriteString("		<ul>")
	for _, review := range view.Reviews {
		html.WriteString("<li>")
		html.WriteString("<div>" + review.StoreName + "</div>")
		html.WriteString("<div>" + review.Title + "</div>")
		html.WriteString("<div>" + review.Date + "</div>")
		if len(review.PictureURLs) > 0 {
			defaultPictureURL := review.PictureURLs[0]
			fmt.Fprintf(html, `<img src=%q></img>`, defaultPictureURL)
		} else {
			html.WriteString(`<p>No picture</p>`)
		}
		fmt.Fprintf(html, `<a href="/reviews/%d">Read more...</a>`, review.ID)
		html.WriteString("</li>")
	}
	html.WriteString("		</ul>")
	html.WriteString("	</body>")
	html.WriteString("</html>")

	_, err := io.Copy(w, html)
	return err
}
