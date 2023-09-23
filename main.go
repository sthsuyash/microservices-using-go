package main

import (
	"fmt"
	"net/http" // this package is used to create a web server

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// fmt.Println("Hello, World!") // Println is a camel case function where first letter is capital because it is public
	// creating a server object

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/hello", basicHandler)

	server := &http.Server{
		Addr: ":8080", // port number
		// Handler: http.HandlerFunc(basicHandler), // handler function // this is for core go
		Handler: router, // this is for chi router
	}

	err := server.ListenAndServe() // ListenAndServe is a blocking function
	if err != nil {
		fmt.Println("Failed to start server", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World from /hello home api!")) // Write function takes a byte array as input

	/*
		This is for core go
	*/
	// if r.Method == http.MethodGet {
	// 	// handle GET
	// 	if r.URL.Path == "/" {
	// 		w.Write([]byte("Hello, World from home api!")) // Write function takes a byte array as input
	// 		return
	// 	}
	// 	if r.URL.Path == "/about" {
	// 		w.Write([]byte("Hello, World from about api!"))
	// 		return
	// 	}
	// }
}
