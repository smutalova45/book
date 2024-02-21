package models

type Book struct {
	ID         string `json:"id"`
	BookName   string `json:"book_name"`
	AuthorName string `json:"author_name"`
	PageNumber int    `json:"page_number"`
}

// update ham uchun Book struct
type CreateBook struct {
	BookName   string `json:"book_name"`
	AuthorName string `json:"author_name"`
	PageNumber int    `json:"page_number"`
}

type UpdatePageNumber struct {
	ID         string `json:"id"`
	PageNumber int    `json:"page_number"`
}

type BookResponse struct {
	Books []Book `json:"books"`
	Count int    `json:"count"`
}

type PrimaryKey struct {
	ID string `json:"id"`
}

type GetListRequest struct {
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}
