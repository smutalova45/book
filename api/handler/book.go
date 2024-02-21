package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"main.go/api/models"
)

// CreateBook godoc
// @Router       /book [POST]
// @Summary      Create a new book
// @Description  create a new book
// @Tags         book
// @Accept       json
// @Produce      json
// @Param 		 book body models.CreateBook true "book"
// @Success      200  {object}  models.Book
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateBook(c *gin.Context) {
	book := models.CreateBook{}
	if err := c.ShouldBindJSON(&book); err != nil {
		handleResponse(c, "error i swhile reading body", http.StatusBadRequest, err.Error())
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	response, err := h.services.Book().Create(ctx, book)
	if err != nil {
		handleResponse(c, "error is while creating book in handler", 500, err.Error())

		return
	}
	handleResponse(c, "", 201, response)
}

// GetBook godoc
// @Router       /book/{id} [GET]
// @Summary      Get book by id
// @Description  get book by id
// @Tags         book
// @Accept       json
// @Produce      json
// @Param 		 id path string true "book_id"
// @Success      200  {object}  models.Book
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetBook(c *gin.Context) {
	uid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	book, err := h.services.Book().GetById(ctx, models.PrimaryKey{ID: uid})
	if err != nil {
		handleResponse(c, "error is getting by id book in handelr", 500, err.Error())
		return
	}
	handleResponse(c, "success", 200, book)

}

// GetListBook godoc
// @Router       /book [GET]
// @Summary      Get book list
// @Description  get book list
// @Tags         book
// @Accept       json
// @Produce      json
// @Param 		 page query string false "page"
// @Param 		 limit query string false "limit"
// @Param 		 search query string false "search"
// @Success      200  {object}  models.BookResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetListBook(c *gin.Context) {
	var (
		page, limit int
		search      string
		err         error
	)
	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c, "error is while converting page", http.StatusBadRequest, err.Error())
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		handleResponse(c, "error is while converting limit", http.StatusBadRequest, err.Error())
		return
	}

	search = c.Query("search")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	books, err := h.services.Book().GetList(ctx, models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		handleResponse(c, "error is while getting book list", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "", http.StatusOK, books)
}

// UpdateBook godoc
// @Router       /book/{id} [PUT]
// @Summary      Update book
// @Description  update book
// @Tags         book
// @Accept       json
// @Produce      json
// @Param 		 id path string true "book_id"
// @Param 		 book body models.Book false "book"
// @Success      200  {object}  models.Book
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateBook(c *gin.Context) {
	uid := c.Param("id")
	book := models.Book{}
	if err := c.ShouldBindJSON(&book); err != nil {
		fmt.Println("ssssssss")
		handleResponse(c, "error while updating book in handelr", 500, err.Error())
		return
	}
	book.ID = uid
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	updatebook, err := h.services.Book().Update(ctx, book)
	if err != nil {
		handleResponse(c, "error is while updating book", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "success", http.StatusOK, updatebook)

}

// DeleteBook godoc
// @Router       /book/{id} [DELETE]
// @Summary      Delete book
// @Description  delete book
// @Tags         book
// @Accept       json
// @Produce      json
// @Param 		 id path string true "book_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteBook(c *gin.Context) {
	uid := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := h.services.Book().Delete(ctx, models.PrimaryKey{ID: uid}); err != nil {
		handleResponse(c, "error is while delting book", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "", 200, "book deleted!")
}

// UpdatePageNumber godoc
// @Router       /book/{id} [PATCH]
// @Summary      Update book
// @Description  update book
// @Tags         book
// @Accept       json
// @Produce      json
// @Param 		 id path string true "book_id"
// @Param 		 book body models.UpdatePageNumber false "book_page_number"
// @Success      200  {object}  models.Book
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdatePageNumber(c *gin.Context) {

	book := models.UpdatePageNumber{}
	if err := c.ShouldBindJSON(&book); err != nil {
		handleResponse(c, "error reading page_number body", 500, err.Error())
		return
	}

	uid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		handleResponse(c, "error while parsing uuid", http.StatusBadRequest, err.Error())
		return
	}

	book.ID = uid.String()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	updatepage, err := h.services.Book().UpdatePageNumber(ctx, book)
	if err != nil {
		handleResponse(c, "error updating page number", 500, err.Error())
		return
	}
	handleResponse(c, "success", 200, updatepage)

}
