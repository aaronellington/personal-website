package assets

import (
	"embed"
	"html/template"
	"io/fs"
)

// Theme for philote
var Theme *template.Template

// Public filesystem for philote
var Public fs.FS

// Content markdown files for philote
var Content fs.FS

//go:embed *
var assets embed.FS

//go:embed theme/theme.go.html
var themeContent string

func init() {
	var err error

	Public, err = fs.Sub(assets, "theme/public")
	if err != nil {
		panic(err)
	}

	Content, err = fs.Sub(assets, "theme/content")
	if err != nil {
		panic(err)
	}

	Theme, err = template.New("theme").Parse(themeContent)
	if err != nil {
		panic(err)
	}
}
