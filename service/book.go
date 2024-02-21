package service

import (
	"context"
	"fmt"

	"main.go/api/models"
	"main.go/storage"
)

type bookService struct {
	storage storage.IStorage
}

func NewBookService(storage storage.IStorage) bookService {
	return bookService{
		storage: storage,
	}
}

func (b bookService) Create(ctx context.Context, book models.CreateBook) (models.Book, error) {
	id, err := b.storage.Book().Create(ctx, book)
	if err != nil {
		fmt.Println("error in service layer while creating book")
		return models.Book{}, err
	}
	createbook, err := b.storage.Book().GetById(context.Background(), models.PrimaryKey{ID: id})
	if err != nil {
		fmt.Println("error getting by id in service layer")
		return models.Book{}, err
	}
	return createbook, err
}

func (b bookService) GetById(ctx context.Context, key models.PrimaryKey) (models.Book, error) {
	book, err := b.storage.Book().GetById(ctx, models.PrimaryKey{ID: key.ID})
	if err != nil {
		fmt.Println("error in service layer while gtting by id of book")
		return models.Book{}, err
	}
	return book, nil

}

func (b bookService) GetList(ctx context.Context, requst models.GetListRequest) (models.BookResponse, error) {
	books, err := b.storage.Book().GetList(ctx, requst)
	if err != nil {
		fmt.Println("error getting list in service layerr")
		return models.BookResponse{}, err
	}
	return books, nil
}

func (b bookService) Update(ctx context.Context, book models.Book) (models.Book, error) {
	id, err := b.storage.Book().Update(ctx, book)
	if err != nil {
		fmt.Println("error while updating book in service layer")
		return models.Book{}, err
	}
	updatedbook, err := b.storage.Book().GetById(context.Background(), models.PrimaryKey{ID: id})
	if err != nil {
		fmt.Println("error while getting list after updating book in service layer")
		return models.Book{}, err
	}
	return updatedbook, nil
}

func (b bookService) Delete(ctx context.Context, key models.PrimaryKey) error {
	err := b.storage.Book().Delete(ctx, key)
	if err != nil {
		fmt.Println("erro in service layer")
		return err
	}
	return nil
}

func (b bookService) UpdatePageNumber(ctx context.Context, pagen models.UpdatePageNumber) (models.Book, error) {
	id, err := b.storage.Book().UpdatePageNumber(ctx, models.UpdatePageNumber{ID: pagen.ID, PageNumber: pagen.PageNumber})
	if err != nil {
		fmt.Println("error while ypdating pagenumber in service layer")
		return models.Book{}, err
	}
	updatepage, err := b.storage.Book().GetById(context.Background(), models.PrimaryKey{ID: id})
	if err != nil {
		fmt.Println("error while getting list after updating book in service layer")
		return models.Book{}, err
	}
	return updatepage, nil

}
