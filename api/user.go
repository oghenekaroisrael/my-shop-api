package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/oghenekaroisrael/myshopapi/db/sqlc"
	"github.com/oghenekaroisrael/myshopapi/utils"
)

// Request Structs
type createUserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

type getUserByIDRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type listUsersRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}

type loginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string       `json:"access_token"`
	User        UserResponse `json:""`
}
type UserResponse struct {
	ID         int32     `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	IsVerified bool      `json:"isVerified"`
	CreatedAt  time.Time `json:"created_at"`
}

func newUserResponse(user db.User) UserResponse {
	return UserResponse{
		ID:         user.ID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		CreatedAt:  user.CreatedAt,
		IsVerified: user.IsVerified,
	}
}

// Controllers
func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	args := db.CreateUserParams{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  hash,
	}

	user, err := server.store.CreateUser(ctx, args)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	resp := newUserResponse(user)
	ctx.JSON(http.StatusCreated, resp)
}

func (server *Server) getUserById(ctx *gin.Context) {
	var req getUserByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUserById(ctx, int32(req.ID))

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	resp := UserResponse{
		ID:         user.ID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		CreatedAt:  user.CreatedAt,
		IsVerified: user.IsVerified,
	}
	ctx.JSON(http.StatusOK, resp)
}

func (server *Server) listUsers(ctx *gin.Context) {
	var req listUsersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.ListUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	users, err := server.store.ListUsers(ctx, args)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := server.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	err = utils.CheckPassword(req.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	// accessToken, err := server.tokenMaker.CreateToken(user.ID, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := LoginResponse{
		User: newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, resp)
}
