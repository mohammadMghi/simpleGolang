package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main()  {
	router:= mux.NewRouter()
	const port string=":8002"
	router.HandleFunc("/" , func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response , "Test")
	})
	router.HandleFunc("/posts" , getPosts).Methods("GET")
	router.HandleFunc("/addPosts" , addPost).Methods("POST")
	log.Println("Server listening on port" , port)
	log.Fatalln(http.ListenAndServe(port,router))

}
