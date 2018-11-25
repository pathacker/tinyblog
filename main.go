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
	<link rel="stylesheet" type="text/css" href="/public/stylesheets/tinyblog.css">
</head>
<body>
	<nav>
		<a href="/">Home</a>
	</nav>
	<article>
	{{.S}}
	<article>
</body>
</html>
`
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var content string
	fn := r.URL.Path[1:]
	if fn == "" {
		fn = "index"
	}
	b, err := ioutil.ReadFile(fmt.Sprintf("content/%s.md", fn))
	if err != nil {
		return
	}
	content = string(blackfriday.Run(b))

	data := struct {
		S template.HTML
	}{
		S: template.HTML(content),
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
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	http.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
