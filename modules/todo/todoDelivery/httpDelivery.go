package todoDelivery

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"todo_service_mongo/models"
	"todo_service_mongo/modules/todo"
)

type httpDelivery struct {
	todoUC todo.Usecase
}

func NewHTTPHandler(todoUC todo.Usecase) {
	handler := httpDelivery{todoUC: todoUC}

	http.HandleFunc("/api/todo", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			handler.getAllTodo(w, req)
		case "POST":
			handler.addNewTodo(w, req)
		case "PUT":
			handler.updateTodo(w, req)
		case "DELETE":
			handler.deleteTodo(w, req)
		default:
			fmt.Fprint(w, "method not supported")
		}
	})
}

func (handler *httpDelivery) getAllTodo(w http.ResponseWriter, req *http.Request) {
	todos, err := handler.todoUC.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonRes, err := json.Marshal(todos)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

type ReqAddNewTodo struct {
	Activity string `json:"activity"`
}

func (handler *httpDelivery) addNewTodo(w http.ResponseWriter, req *http.Request) {
	var reqBody ReqAddNewTodo
	err := json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newTodo, err := handler.todoUC.Insert(reqBody.Activity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonRes, err := json.Marshal(newTodo)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

type ReqUpdateTodo struct {
	DocumentID string `json:"document_id"`
	Activity   string `json:"activity"`
}

func (handler *httpDelivery) updateTodo(w http.ResponseWriter, req *http.Request) {
	var reqBody ReqUpdateTodo
	err := json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//updated todo
	objectId, err := primitive.ObjectIDFromHex(reqBody.DocumentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newTodo := models.Todo{Activity: reqBody.Activity, DocumentID: objectId}

	updatedTodo, err := handler.todoUC.Update(newTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonRes, err := json.Marshal(updatedTodo)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

type ReqDeleteTodo struct {
	DocumentID string `json:"document_id"`
}

func (handler *httpDelivery) deleteTodo(w http.ResponseWriter, req *http.Request) {
	var reqBody ReqDeleteTodo
	err := json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.todoUC.DeleteByID(reqBody.DocumentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var res struct {
		Message string `json:"message"`
 	}

	res.Message = reqBody.DocumentID + " is deleted successfully"

	jsonRes, err := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}
