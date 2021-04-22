package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"
)

var cs *sessions.CookieStore = sessions.NewCookieStore([]byte("secret-key-1234"))

// type Temps struct {
// 	notemp *template.Template
// 	indx   *template.Template
// 	helo   *template.Template
// }

func notemp() *template.Template {
	src := "<html><body><h1>NO TEMPLATE.</h1></body></html>"
	tmp, _ := template.New("index").Parse(src)
	return tmp
}

func page(fname string) *template.Template {
	tmps, _ := template.ParseFiles("templates/"+fname+".html", "templates/head.html", "templates/foot.html")
	return tmps

}

// func setupTemp() *Temps {
// 	temps := new(Temps)

// 	temps.notemp = notemp()

// 	indx, er := template.ParseFiles("templates/index.html")

// 	if er != nil {
// 		indx = temps.notemp

// 	}
// 	temps.indx = indx

// 	helo, er := template.ParseFiles("templates/hello.html")

// 	if er != nil {
// 		helo = temps.notemp
// 	}
// 	temps.helo = helo

// 	return temps

// }

func index(w http.ResponseWriter, rq *http.Request) {
	item := struct {
		Template string
		Title    string
		Message  string
	}{
		Template: "index",
		Title:    "Index",
		Message:  "This is Go Top page",
	}

	er := page("index").Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}

// var cs *sessions.CookieStore = sessions.NewCookieStore([]byte("secret-key-12345"))

func hello(w http.ResponseWriter, rq *http.Request) {

	data := []string{
		"One", "Two", "Three",
	}

	item := struct {
		Title string
		Data  []string
	}{
		Title: "Hello",
		Data:  data,
	}

	er := page("hello").Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}

}

func main() {

	// temps := setupTemp()

	http.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
		index(w, rq)

	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, rq *http.Request) {
		hello(w, rq)
	})

	http.ListenAndServe(":8000", nil)

}
