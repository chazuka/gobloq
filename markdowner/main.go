package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/russross/blackfriday"
	"html/template"
)

func parseMarkdown(args ...interface{}) template.HTML {
	s := blackfriday.MarkdownCommon([]byte(fmt.Sprintf("%s", args...)))
	return template.HTML(s)
}

func main() {

	engine := martini.Classic()

	engine.Use(render.Renderer(render.Options{
		Directory:  "views",
		Layout:     "master",
		Funcs:      []template.FuncMap{{"pmd": parseMarkdown}},
		Extensions: []string{".md", ".html"},
		Charset:    "UTF-8",
	}))

	engine.Get("/", func(r render.Render) {
		r.HTML(200, "homepage", map[string]interface{}{"Title": "Golangers"})
	})

	engine.Run()

}
