package api

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	db "github.com/naviscom/catalystx2/db/sqlc"
)

type createUserRequest struct {
	Username	string	`json:"username" binding:"required,alphanum"`
	Password	string	`json:"password" binding:"required,min=6"`
	FullName	string	`json:"full_name" binding:"required"`
	Email	string	`json:"email" binding:"required,email"`
}

type userResponse struct {
	Username	string	`json:"username`
	FullName	string	`json:"full_name`
	Email	string	`json:"email`
	PasswordChangedAt	string	`json:"password_changed_at`
	PasswordCreatedAt	string	`json:"password_created_at`
	}

func newUserResponse(user db.User) userResponse {
	return userResponse {
	Username: user.Username,
	FullName: user.FullName,
	Email: user.Email,
	PasswordChangedAt: user.PasswordChangedAt,
	PasswordCreatedAt: user.PasswordCreatedAt,
	}
}
func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:	req.Username,
		HashedPassword:	hashedPassword,
		FullName:	req.FullName,
		Email:	req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	resp := newUserResponse(user)
	ctx.JSON(http.StatusOK, resp)
}

type loginUserRequest struct {
	Username	string	`json:"username" binding:"required,alphanum"`
	Password	string	`json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken	string	`json:"access_token"`	User	userResponse	`json:"user"`}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		user.Username,
//		user.Role,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

//	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
//		user.Username,
//		user.Role,
//		server.config.RefreshTokenDuration,
//	)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//		return
//	}

//	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
//		ID:           refreshPayload.ID,
//		Username:     user.Username,
//		RefreshToken: refreshToken,
//		UserAgent:    ctx.Request.UserAgent(),
//		UserAgent:    ctx.Request.UserAgent(),
//		IsBlocked:    false,
//		ExpiresAt:    refreshPayload.ExpiredAt,
//	})
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//		return
//	}

	rsp := loginUserResponse{
//		SessionID:             session.ID,
		AccessToken:           accessToken,
//		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
//		RefreshToken:          refreshToken,
//		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, rsp)

type getUserRequest0 struct {
	Username	string	`uri:"username" binding:"required,min=1"`
}

func (server *Server) getUser0(ctx *gin.Context) {
	var req getUserRequest0
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser0(ctx, req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type getUserRequest3 struct {
	Email	string	`uri:"email" binding:"required,min=1"`
}

func (server *Server) getUser3(ctx *gin.Context) {
	var req getUserRequest3
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser3(ctx, req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type listUserRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize   int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listUsers(ctx *gin.Context) {
	var req listUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	users, err := server.store.ListUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, users)
}

type updateUserRequest struct {
	Username	string	`json:"username" binding:"required"`
	HashedPassword	string	`json:"hashed_password" binding:"required"`
	FullName	string	`json:"full_name" binding:"required"`
	Email	string	`json:"email" binding:"required"`
	PasswordChangedAt	time.Time	`json:"password_changed_at" binding:"required"`
	PasswordCreatedAt	time.Time	`json:"password_created_at" binding:"required"`
}

func (server *Server) updateUser(ctx *gin.Context) {
	var req updateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateUserParams{
		Username:	req.Username,
		HashedPassword:	req.HashedPassword,
		FullName:	req.FullName,
		Email:	req.Email,
		PasswordChangedAt:	req.PasswordChangedAt,
	}
	user, err := server.store.UpdateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type deleteUserRequest struct {
	Username	string	`uri:"username" binding:"required,alphanum"`
}

func (server *Server) deleteUser(ctx *gin.Context) {
		var req deleteUserRequest
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		err := server.store.DeleteUser(ctx, req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "record deleted successfully")
}

