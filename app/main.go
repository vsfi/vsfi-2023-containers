package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

var port = flag.Int("port", 8080, "Specify the port to listen to.")

var db, _ = IntDatabase()

type Page struct {
	Name     string
	Response string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var name, _ = os.Hostname()
	tmpl := template.Must(template.ParseFiles("templates/create.html"))
	tmpl.Execute(w, name)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	count, _ := strconv.Atoi(r.FormValue("count"))
	name := r.FormValue("name")

	err := AddToQueue(db, name, count)
	if err != nil {
		log.Println(err)
	}

	p := Page{Name: name, Response: "Your ticket has been accepted! Spasibo for choose our service!"}

	tmpl := template.Must(template.ParseFiles("templates/create.html"))
	tmpl.Execute(w, p)
}

func queueHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	data := GetQueue(db, name)
	tmpl := template.Must(template.ParseFiles("templates/queue.html"))
	tmpl.Execute(w, data)
}

func main() {
	fmt.Fprintf(os.Stdout, "Web Server started. Listening on 0.0.0.0:%d\n", *port)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/queue", queueHandler)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		fmt.Println(err.Error())
	}
}
