package blogrenderer

import "io"

// create a post struct

type Post struct {
	Title       string
	Body        string
	Description string
	Tags        []string
}

// create method Render with io.Writer

func Render(w io.Writer, p Post) error {
	return nil
}
