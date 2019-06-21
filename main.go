package main

import (
	"fmt"
	"github.com/husobee/vestigo"
	"net/http"
)

func main() {
	router := vestigo.NewRouter()
	router.Get("/*", http.FileServer(http.Dir("web/build")).ServeHTTP)

	vestigo.CustomNotFoundHandlerFunc(customHandler)

	http.ListenAndServe(":1234", router)
}

func customHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("customHandler")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("cccccccccccccccccustom not found"))

}
