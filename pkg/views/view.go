package views

import "io"

type View interface {
	ContentType() string
	Render(io.Writer) error
}
