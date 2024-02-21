package storage

import (
	"context"

	"main.go/api/models"
)

type IStorage interface {
	Close()
	Book() IBookStorage
}
type IBookStorage interface {
	Create(context.Context, models.CreateBook) (string, error)
	GetById(context.Context, models.PrimaryKey) (models.Book, error)
	GetList(context.Context, models.GetListRequest) (models.BookResponse, error)
	Update(context.Context, models.Book) (string, error)
	Delete(context.Context, models.PrimaryKey) error
	UpdatePageNumber(context.Context, models.UpdatePageNumber) (string, error)
}
