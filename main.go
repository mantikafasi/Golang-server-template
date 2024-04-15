package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	"server-go/common"
	"server-go/routes"
)


func main() {

	common.InitCache()
	// uncomment if you have a database set up
	// database.InitDB()


	mux := chi.NewRouter()
	
	mux.Use(routes.CorsMiddleware)

	mux.Handle("/", http.FileServer(http.Dir("website")))

	mux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "An Error occurred\n")
	})

	fmt.Printf("server started at http://localhost:%s\n", common.Config.Port)

	err := http.ListenAndServe(":"+common.Config.Port, mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")

	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
