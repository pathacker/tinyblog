/*
   Tinyblog is a super simple web blog.
*/

package main

import (
	"html/template"
	"log"
	"net/http"
)

var (
	tmpl template.Template
	s    = `
<!DOCTYPE html>
<head>
	<link rel="stylesheet" href="/public/stylesheets/tinyblog.css" type="text/css">
</head>
<body>
	<h1>
		{{.S}}
	</h1>
	<p>Foobar</p>
</body>
</html>
`
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		S string
	}{
		S: "found me!",
	}
	t := template.Must(template.New("foo").Parse(s))
	err := t.Execute(w, data)
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
