package assets

import (
	"embed"
	"html/template"
	"io/fs"
)

var Theme *template.Template
var Public fs.FS
var Content fs.FS

//go:embed *
var Assets embed.FS

//go:embed theme/theme.go.html
var themeContent string

func init() {
	var err error

	Public, err = fs.Sub(Assets, "theme/public")
	if err != nil {
		panic(err)
	}

	Content, err = fs.Sub(Assets, "theme/content")
	if err != nil {
		panic(err)
	}

	Theme, err = template.New("theme").Parse(themeContent)
	if err != nil {
		panic(err)
	}
}
