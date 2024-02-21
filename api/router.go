package api

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"main.go/api/handler"
	"main.go/service"
	_ "main.go/api/docs"
)

// New ...
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @security definitions.apikey
func New(services service.IServiceManager) *gin.Engine {
	h := handler.New(services)
	r := gin.New()
	r.Use(traceRequest)
	{
		r.POST("/book", h.CreateBook)
		r.GET("/book/:id", h.GetBook)
		r.GET("/book", h.GetListBook)
		r.PUT("/book/:id", h.UpdateBook)
		r.DELETE("/book/:id", h.DeleteBook)
		r.PATCH("/book/:id", h.UpdatePageNumber)

		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	}
	return r
}

func traceRequest(c *gin.Context) {
	beforeRequest(c)
	c.Next()
	afterRequest(c)

}
func beforeRequest(c *gin.Context) {
	startTime := time.Now()
	c.Set("start_time", startTime)

	log.Println("start time:", startTime.Format("2006-01-02 15:04:05.0000"), "path:", c.Request.URL.Path)

}
func afterRequest(c *gin.Context) {
	startTime, exists := c.Get("start_time")
	if !exists {
		startTime = time.Now()
	}

	duration := time.Since(startTime.(time.Time)).Milliseconds()

	log.Printf("%s %s %s status:%d, time:%d milliseconds\n",
		startTime.(time.Time).Format("2006-01-02 15:04:05.000"), c.Request.Method, c.Request.URL.Path, c.Writer.Status(), duration)
}
