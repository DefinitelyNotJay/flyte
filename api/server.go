package api

import (
	"fmt"

	db "github.com/DefinitelyNotJay/flyte/db/sqlc"
	"github.com/DefinitelyNotJay/flyte/token"
	"github.com/DefinitelyNotJay/flyte/util"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config util.Config
	store  db.Store
	tokenMaker token.Maker
	router *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config: config,
		store: store,
		tokenMaker: tokenMaker,
	}
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.GET("/users", server.listUser)
	router.GET("/users/:id", server.getUser)

	server.router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
