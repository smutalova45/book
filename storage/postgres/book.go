package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"main.go/api/models"
	"main.go/storage"
)

type bookRepo struct {
	db *pgxpool.Pool
}

func NewBookRepo(db *pgxpool.Pool) storage.IBookStorage {
	return &bookRepo{
		db: db,
	}
}
func (b *bookRepo) Create(ctx context.Context, book models.CreateBook) (string, error) {
	id := uuid.New()
	query := `insert into book(id, book_name, author_name, page_number)
	            values($1, $2, $3, $4)`
	if _, err := b.db.Exec(ctx, query, id, book.BookName, book.AuthorName, book.PageNumber); err != nil {
		fmt.Println("error while creating book in handler:", err)
		return "", err
	}

	return id.String(), nil
}

func (b *bookRepo) GetById(ctx context.Context, key models.PrimaryKey) (models.Book, error) {
	book := models.Book{}
	query := `select id, book_name, author_name, page_number from book where id =$1 `
	if err := b.db.QueryRow(ctx, query, key.ID).Scan(
		&book.ID,
		&book.BookName,
		&book.AuthorName,
		&book.PageNumber); err != nil {
		return models.Book{}, err
	}

	return book, nil
}
func (b *bookRepo) GetList(ctx context.Context, request models.GetListRequest) (models.BookResponse, error) {
	var (
		books  = []models.Book{}
		page   = request.Page
		offset = (page - 1) * request.Limit
		search = request.Search
		query  string
		count  = 0
	)
	countquery := `select count(1) from book `
	if search != "" {
		countquery += fmt.Sprintf(` where book_name ilike '%%%s%%'`, search)
	}
	if err := b.db.QueryRow(ctx, countquery).Scan(&count); err != nil {
		fmt.Println("error getting list of books in postgres")
		return models.BookResponse{}, err
	}

	query = `select id, book_name, author_name, page_number from book`

	if search != "" {
		query += fmt.Sprintf(` where book_name ilike '%%%s%%'`, search)
	}
	query += ` LIMIT $1 OFFSET $2`

	rows, err := b.db.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error while executing query:", err)
		return models.BookResponse{}, err
	}
	for rows.Next() {
		book := models.Book{}
		if err = rows.Scan(&book.ID, &book.BookName, &book.AuthorName, &book.PageNumber); err != nil {
			fmt.Println("error while getting list of books in postgres")
			return models.BookResponse{}, err
		}

		books = append(books, book)
	}

	return models.BookResponse{
		Books: books,
		Count: count,
	}, nil
}

func (b *bookRepo) Update(ctx context.Context, book models.Book) (string, error) {
	query := `UPDATE book SET book_name=$1, author_name=$2, page_number=$3 WHERE id=$4`
	if _, err := b.db.Exec(ctx, query, &book.BookName, &book.AuthorName, &book.PageNumber, &book.ID); err != nil {
		fmt.Println("error updating book:", err)
		return "", err
	}

	return book.ID, nil
}


func (b *bookRepo) Delete(ctx context.Context, key models.PrimaryKey) error {
	query := ` delete from book where id=$1`
	fmt.Println("lllll")
	if _, err := b.db.Exec(ctx, query, key.ID); err != nil {
		fmt.Println("error!!!")
		return err
	}

	return nil

}

func (b *bookRepo) UpdatePageNumber(ctx context.Context, pagenumber models.UpdatePageNumber) (string, error) {
	query := ` update book set page_number=$1 where id=$2`
	if rowsAffected, err := b.db.Exec(ctx, query, pagenumber.PageNumber, pagenumber.ID); err != nil {
		if r := rowsAffected.RowsAffected(); r == 0 {
			fmt.Println("no rows were affected by the update operation")
			return "", err
		}
		return "", err
	}
	return pagenumber.ID, nil
}
