package model

import "github.com/a-h/templ"

type PostData struct {
	Slug        string
	Title       string `toml:"title"`
	Description string `toml:"description"`
	Author      struct {
		Name  string `toml:"name"`
		Email string `toml:"email"`
	} `toml:"author"`
}

type Post struct {
	PostData
	Content templ.Component
}
