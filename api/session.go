package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/naviscom/catalystx2/db/sqlc"
)

type createSessionRequest struct {
	Username     string    `json:"username" binding:"required"`
	RefreshToken string    `json:"refresh_token" binding:"required"`
	UserAgent    string    `json:"user_agent" binding:"required"`
	ClientIp     string    `json:"client_ip" binding:"required"`
	ExpiresAt    time.Time `json:"expires_at" binding:"required"`
	CreatedAt    time.Time `json:"created_at" binding:"required"`
}

func (server *Server) createSession(ctx *gin.Context) {
	var req createSessionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateSessionParams{
		ID:           req.ID,
		Username:     req.Username,
		RefreshToken: req.RefreshToken,
		UserAgent:    req.UserAgent,
		ClientIp:     req.ClientIp,
		IsBlocked:    req.IsBlocked,
		ExpiresAt:    req.ExpiresAt,
		CreatedAt:    req.CreatedAt,
	}
	session, err := server.store.CreateSession(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, session)
}

type getSessionRequest0 struct {
	ID `uri:"id" binding:"required,min=1"`
}

func (server *Server) getSession0(ctx *gin.Context) {
	var req getSessionRequest0
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	session, err := server.store.GetSession0(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, session)
}

type listSessionRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listSessions(ctx *gin.Context) {
	var req listSessionRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListSessionsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	sessions, err := server.store.ListSessions(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, sessions)
}

type updateSessionRequest struct {
	ID           `json:"id" binding:"required"`
	Username     string `json:"username" binding:"required"`
	RefreshToken string `json:"refresh_token" binding:"required"`
	UserAgent    string `json:"user_agent" binding:"required"`
	ClientIp     string `json:"client_ip" binding:"required"`
	IsBlocked    `json:"is_blocked" binding:"required"`
	ExpiresAt    time.Time `json:"expires_at" binding:"required"`
	CreatedAt    time.Time `json:"created_at" binding:"required"`
}

func (server *Server) updateSession(ctx *gin.Context) {
	var req updateSessionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateSessionParams{
		ID:           req.ID,
		Username:     req.Username,
		RefreshToken: req.RefreshToken,
		UserAgent:    req.UserAgent,
		ClientIp:     req.ClientIp,
		IsBlocked:    req.IsBlocked,
		ExpiresAt:    req.ExpiresAt,
		CreatedAt:    req.CreatedAt,
	}
	session, err := server.store.UpdateSession(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, session)
}

type deleteSessionRequest struct {
	ID uuid.UUID `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteSession(ctx *gin.Context) {
	var req deleteSessionRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteSession(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "record deleted successfully")
}
