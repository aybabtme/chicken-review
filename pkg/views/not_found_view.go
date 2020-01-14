package views

import "io"

type NotFoundView struct{}

func (view NotFoundView) ContentType() string {
	return "text/html"
}

func (view NotFoundView) Render(w io.Writer) error {
	notFoundHTML := `
<html>
	<body>
		<h1>Sorry, not found!</h1>
	</body>
</html>
	`
	_, err := w.Write([]byte(notFoundHTML))
	return err
}
