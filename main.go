package main

import (
	"fmt"
	"log"
	"net/http"
)

type path struct {
	address map[string]string
}

var root = initRooter().address

func initRooter() *path {
	r := path{
		address: map[string]string{
			"/":     "/",
			"hello": "/hello",
			"form":  "/form",
		},
	}

	return &r
}

func server() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle(root["/"], fileServer)
	http.HandleFunc(root["form"], formHandle)
	http.HandleFunc(root["hello"], helloHandle)
}

func cekPath(path string, method string, w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != path {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != method {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}

func helloHandle(w http.ResponseWriter, r *http.Request) {
	cekPath(r.URL.Path, "GET", w, r)
	fmt.Fprintf(w, "Hello")
}

func checkFormPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error: %v", err)
		return
	}
}

func formHandle(w http.ResponseWriter, r *http.Request) {
	cekPath(r.URL.Path, "POST", w, r)
	checkFormPost(w, r)

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "name : %s\naddress : %s\n", name, address)
}

func runServer() {
	server()
	fmt.Println("Run server localhost:9090")

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	runServer()
}
