package todo

import "todo_service_mongo/models"

type Repository interface{
	FindAll()(*[]models.Todo, error)
	Insert(todo models.Todo)(*models.Todo, error)
	Update(todo models.Todo)(*models.Todo, error)
	DeleteByID(id string) error
}
