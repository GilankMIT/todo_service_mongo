package todoRepository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
	"todo_service_mongo/models"
	"todo_service_mongo/modules/todo"
)

type mongoRepo struct{
	mongoDB *mongo.Database
}

//NewMongoRepo return instance implementing todo.Repository in MongoDB
func NewMongoRepo(mongoClient *mongo.Database) todo.Repository{
	return &mongoRepo{mongoDB: mongoClient}
}

//FindAll return all todo in mongoDB
func (m *mongoRepo) FindAll() (*[]models.Todo, error) {
	var todos []models.Todo

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := m.mongoDB.Collection("todo").Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx){
		var todo models.Todo
		err = cur.Decode(&todo)
		if err != nil{
			continue
		}
		todos = append(todos, todo)
	}

	return &todos, nil
}


//Insert will add new document to MongoDB
func (m *mongoRepo) Insert(todo models.Todo) (*models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := m.mongoDB.Collection("todo").InsertOne(ctx, bson.M{"activity" : todo.Activity})
	if err != nil{
		return nil, err
	}

	//Add last inserted document id from MongoDB to the inserted data
	todo.DocumentID = res.InsertedID.(primitive.ObjectID)
	return &todo, nil
}

//Update will modify / alter existing document based on filter (ObjectID)
func (m *mongoRepo) Update(todo models.Todo) (*models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := m.mongoDB.Collection("todo").UpdateOne(ctx,
		bson.M{"_id" : todo.DocumentID}, bson.M{"$set" : &todo})
	if err != nil{
		return nil, err
	}
	return &todo, nil
}

//DeleteByID delete MongoDB todo document by ID
func (m *mongoRepo) DeleteByID(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Parse string to ObjectID
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return err
	}

	_, err = m.mongoDB.Collection("todo").
		DeleteOne(ctx, bson.M{"_id" : objectId})
	if err != nil{
		return err
	}
	return nil
}
