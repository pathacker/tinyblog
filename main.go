/*
   Tinyblog is a super simple web blog.
*/

package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

var (
	tmpl       template.Template
	tmplString = `
<!DOCTYPE html>
<head>
	<style>{{.Stylesheet}}</style>
</head>
<body>
	<article>
	{{.S}}
	<article>
</body>
</html>
`
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var s string
	fn := r.URL.Path[1:]
	ssbuf, err := ioutil.ReadFile("public/stylesheets/tinyblog.css")
	stylesheet := string(ssbuf)
	if err != nil {
		log.Fatal(err)
	}
	if fn == "" {
		b, err := ioutil.ReadFile("content/index.html")
		s = string(b)
		if err != nil {
			return
		}
	} else {
		b, err := ioutil.ReadFile(fmt.Sprintf("content/%s.md", fn))
		s = string(blackfriday.Run(b))
		if err != nil {
			return
		}
	}

	data := struct {
		S          template.HTML
		Stylesheet template.CSS
	}{
		S:          template.HTML(s),
		Stylesheet: template.CSS(stylesheet),
	}
	t := template.Must(template.New("foo").Parse(tmplString))
	err = t.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public", fs))
	http.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
