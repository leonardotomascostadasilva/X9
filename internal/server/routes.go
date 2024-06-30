package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.GET("/", Handle)

	return r
}

func Handle(c *gin.Context) {

	c.JSON(http.StatusCreated, "Hello")
}
