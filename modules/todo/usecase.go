package todo

import "todo_service_mongo/models"

type Usecase interface{
	GetAll()(*[]models.Todo, error)
	Insert(activity string)(*models.Todo, error)
	Update(todo models.Todo)(*models.Todo, error)
	DeleteByID(id string) error
}
