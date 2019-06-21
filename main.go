package main

import (
	"github.com/husobee/vestigo"
	"net/http"
)

func main() {
	router := vestigo.NewRouter()

	router.Get("/*", http.FileServer(http.Dir("static")).ServeHTTP)

	http.ListenAndServe(":1234", router)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	name := vestigo.Param(r, "name")
	w.Write([]byte("Hey " + name))
}
