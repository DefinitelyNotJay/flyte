package api

import (
	db "flyte/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type createUserRequest struct {
	Username     string      `json:"username"`
	Email        string      `json:"email"`
	PasswordHash string      `json:"password_hash"`
	DisplayName  pgtype.Text `json:"display_name"`
	Bio          pgtype.Text `json:"bio"`
	AvatarUrl    pgtype.Text `json:"avatar_url"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: req.PasswordHash,
		DisplayName:  req.DisplayName,
		Bio:          req.Bio,
		AvatarUrl:    req.AvatarUrl,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}
