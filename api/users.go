package api

import (
	"database/sql"
	"net/http"

	db "github.com/Imlucky883/simple_bank/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Username     string `json:"username" binding:"required,alphanum"`
	HashPassword string `json:"hash_password" binding:"required,min=6"`
	FullName     string `json:"full_name" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
}

type createUserResponse struct {
	Username        string `json:"username"`
	FullName        string `json:"full_name"`
	Email           string `json:"email"`
	PasswordChanged string `json:"password_changed"`
	CreatedAt       string `json:"created_at"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:     req.Username,
		HashPassword: req.HashPassword,
		FullName:     req.FullName,
		Email:        req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := createUserResponse{
		Username:        user.Username,
		FullName:        user.FullName,
		Email:           user.Email,
		PasswordChanged: user.PasswordChanged.String(),
		CreatedAt:       user.CreatedAt.String(),
	}

	ctx.JSON(http.StatusOK, resp)
}

type getUserRequest struct {
	Username string `uri:"username" binding:"required,alphanum"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
