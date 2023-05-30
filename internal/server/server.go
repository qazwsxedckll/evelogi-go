package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	*gin.Engine
}

func NewServer() *Server {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return &Server{r}
}

func (s *Server) Run(addr string) error {
	return s.Engine.Run(addr)
}
