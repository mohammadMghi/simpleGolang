package main

import (
	"./entity"
	"encoding/json"
	"math/rand"
	"net/http"
	"./repository"
)


var (
	repo = repository.NewPostRepository()
)


func getPosts(response http.ResponseWriter, request *http.Request)  {
	response.Header().Set("Content-type" , "application/json")
	posts , error:=repo.FindAll()
	if error != nil{
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{error :"Error getting posts"}`))
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}



func addPost(response http.ResponseWriter, request *http.Request)  {
	var post entity.Post
	response.Header().Set("Content-type" , "application/json")
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil{
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{error :"Error Marshaling the posts array"}`))
		return
	}
	post.ID = rand.Int()
	repo.Save(&post)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(post)
}