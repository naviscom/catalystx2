package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/naviscom/catalystx2/db/sqlc"

	/////////////////////////////////////////
	"errors"

	"github.com/naviscom/catalystx2/token"
	/////////////////////////////////////////
)

type createClutterRequest struct {
	ClutterName string `json:"clutter_name" binding:"required"`
	ClutterDesc string `json:"clutter_desc" binding:"required"`
}

func (server *Server) createClutter(ctx *gin.Context) {
	////////////////////////////////////////////////////////////////////////
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authPayload.Role == "level_1_user" {
		err := errors.New("user is not authorized to perform this activity")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	////////////////////////////////////////////////////////////////////////
	var req createClutterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateClutterParams{
		ClutterName: req.ClutterName,
		ClutterDesc: req.ClutterDesc,
	}
	clutter, err := server.store.CreateClutter(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, clutter)
}

type getClutterRequest0 struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getClutter0(ctx *gin.Context) {
	var req getClutterRequest0
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	clutter, err := server.store.GetClutter0(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, clutter)
}

type getClutterRequest1 struct {
	ClutterName string `uri:"clutter_name" binding:"required,min=1"`
}

func (server *Server) getClutter1(ctx *gin.Context) {
	var req getClutterRequest1
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	clutter, err := server.store.GetClutter1(ctx, req.ClutterName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, clutter)
}

type listClutterRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listClutters(ctx *gin.Context) {
	var req listClutterRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListCluttersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	clutters, err := server.store.ListClutters(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, clutters)
}

type updateClutterRequest struct {
	ID          int64  `json:"id" binding:"required"`
	ClutterName string `json:"clutter_name" binding:"required"`
	ClutterDesc string `json:"clutter_desc" binding:"required"`
}

func (server *Server) updateClutter(ctx *gin.Context) {
	////////////////////////////////////////////////////////////////////////
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authPayload.Role == "level_1_user" {
		err := errors.New("user is not authorized to perform this activity")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	////////////////////////////////////////////////////////////////////////
	var req updateClutterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateClutterParams{
		ID:          req.ID,
		ClutterName: req.ClutterName,
		ClutterDesc: req.ClutterDesc,
	}
	clutter, err := server.store.UpdateClutter(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, clutter)
}

type deleteClutterRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteClutter(ctx *gin.Context) {
	////////////////////////////////////////////////////////////////////////
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authPayload.Role == "level_1_user" {
		err := errors.New("user is not authorized to perform this activity")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	////////////////////////////////////////////////////////////////////////
	var req deleteClutterRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteClutter(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "record deleted successfully")
}
