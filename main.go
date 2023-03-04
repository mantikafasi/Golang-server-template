package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"server-go/common"
	"server-go/database"
)

type Cors struct {
	handler *http.ServeMux
}

func (c *Cors) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	c.handler.ServeHTTP(w, r)
}

func (c *Cors) HandleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	c.handler.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	})
}

func (c *Cors) Handle(pattern string, handler http.Handler) {
	c.handler.Handle(pattern, handler)
}

func main() {

	common.InitCache()
	database.InitDB()

	mux := &Cors{http.NewServeMux()}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "website/index.html")
	})

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("website/static"))))
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("website/assets"))))

	mux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "An Error occurred\n")
	})

	err := http.ListenAndServe(":"+common.Config.Port, mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")

	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
