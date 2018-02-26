/*
   Tinyblog is a super simple web blog.
*/

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var (
	tmpl template.Template
	s    = `
<!DOCTYPE html>
<head>
</head>
<body>
<h1>
{{.S}}
</h1>
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
	fmt.Printf("%s\n", r.URL.Path)
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	http.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
