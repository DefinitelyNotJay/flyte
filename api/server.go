package api

import (
	db "github.com/DefinitelyNotJay/flyte/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func Newserver(store db.Store) *Server {
	server := &Server{
		store: store,
	}
	router := gin.Default()

	users := router.Group("/users")

	users.POST("/", server.createUser)
	users.GET("/", server.listUser)
	users.GET("/:id", server.getUser)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
