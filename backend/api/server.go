package api

import (
	"fmt"

	db "github.com/DefinitelyNotJay/flyte/db/sqlc"
	"github.com/DefinitelyNotJay/flyte/token"
	"github.com/DefinitelyNotJay/flyte/util"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	router.Use(cors.New(config))

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/users", server.listUser)
	authRoutes.GET("/users/:id", server.getUser)

	authRoutes.POST("/posts", server.createPost)
	authRoutes.GET("/posts/:id", server.getPost)
	authRoutes.GET("/posts", server.listPosts)
	authRoutes.PATCH("/posts", server.updatePost)
	authRoutes.DELETE("/posts/:id", server.deletePost)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
