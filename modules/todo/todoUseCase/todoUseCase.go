package todoUseCase

import (
	"todo_service_mongo/models"
	"todo_service_mongo/modules/todo"
)

type usecase struct{
	todoRepo todo.Repository
}


//NewUseCase return instance implementing todo.Usecase
func NewUsecase(todoRepo todo.Repository) todo.Usecase{
	return &usecase{todoRepo: todoRepo}
}

//GetAll return all todo data from repository
func (u *usecase) GetAll() (*[]models.Todo, error) {
	return u.todoRepo.FindAll()
}

//Insert will add new data to repository
func (u *usecase) Insert(activity string) (*models.Todo, error) {
	todo := models.Todo{Activity: activity}
	newTodo, err := u.todoRepo.Insert(todo)
	if err != nil{
		return nil, err
	}

	return newTodo, nil
}

//Update will modify existing data in repository
func (u *usecase) Update(todo models.Todo) (*models.Todo, error) {
	newTodo, err := u.todoRepo.Update(todo)
	if err != nil{
		return nil, err
	}

	return newTodo, nil
}

//DeleteByID will remove data in repository based on it's ID / Primary Key
func (u *usecase) DeleteByID(id string) error {
	return u.todoRepo.DeleteByID(id)
}
