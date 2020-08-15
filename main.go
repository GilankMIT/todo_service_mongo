package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
	"todo_service_mongo/modules/todo/todoDelivery"
	"todo_service_mongo/modules/todo/todoRepository"
	"todo_service_mongo/modules/todo/todoUseCase"
)


func main(){
	//MongoDB Host URI
	mongoDBURI := "mongodb://localhost:27017"

	//define MongoDB credential
	var cred options.Credential
	cred.Username = "admin"
	cred.Password = "root"

	//Define connection option to MongoDB by applying URI and Auth credential
	var mongoClientOption options.ClientOptions
	mongoClientOption.ApplyURI(mongoDBURI)
	mongoClientOption.SetAuth(cred)
	mongoClientOption.SetMaxPoolSize(100)


	//Set client connection timeout to 10 second
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	//initialize mongoDB client connection
	client, err := mongo.Connect(ctx, &mongoClientOption)
	if err != nil{
		panic(err)
	}

	//initialize mongoDB client database connection
	mongoDB := client.Database("todo_service")


	//initialize domain modules
	todoRepo := todoRepository.NewMongoRepo(mongoDB)
	todoUC := todoUseCase.NewUsecase(todoRepo)
	todoDelivery.NewHTTPHandler(todoUC)

	//start HTTP server
	log.Println("Server start on localhost:4021")
	if err = http.ListenAndServe(":4021", nil); err != nil{
		panic(err)
	}
}
