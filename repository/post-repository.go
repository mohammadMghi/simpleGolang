package repository

import (
	"../entity"
    "context"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"time"
)
type PostRepository interface {
	Save(post *entity.Post)(interface{},error)
	FindAll()([]entity.Post,error)
}

type repo struct {}

func NewPostRepository() PostRepository {
	return &repo{}
}




func (*repo) Save(post *entity.Post )(interface{},error){
	ctx , _ :=context.WithTimeout(context.Background(),10*time.Second)
	client, _ := mongo.Connect(ctx,options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("myDatabase").Collection("posts")
	result , _ :=collection.InsertOne(ctx,post)
	return result.InsertedID  , nil
}

var arrPosts []entity.Post
func (*repo)FindAll()([]entity.Post,error)  {
	ctx , _ :=context.WithTimeout(context.Background(),10*time.Second)
	client, _ := mongo.Connect(ctx,options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("myDatabase").Collection("posts")
	cursor, _ := collection.Find(context.TODO(), bson.D{})
	var posts []entity.Post
	for cursor.Next(ctx) {
		var post entity.Post
		if err := cursor.Decode(&post); err != nil{
			log.Fatal(err)
		}
		arrPosts =append(posts,post)

	}
	return arrPosts  , nil
}