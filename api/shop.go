package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/oghenekaroisrael/myshopapi/db/sqlc"
)

// Request Structs
type createShopRequest struct {
	ShopName string `json:"shop_name" binding:"required"`
	ShopType string `json:"shop_type" binding:"required"`
	Address  string `json:"address" binding:"required"`
	UserID   int32  `json:"user_id" binding:"required"`
}

type getShopByIDRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type deleteShopByIDRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type listMyShopsRequest struct {
	UserID   int32 `form:"user_id" binding:"required"`
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}

type UpdateShopDetailRequest struct {
	ID       int32  `json:"id" binding:"required"`
	ShopName string `json:"shop_name" binding:"required"`
	ShopType string `json:"shop_type" binding:"required"`
	Address  string `json:"address" binding:"required"`
}

type StatResponse struct {
	Count int64 `json:"count"`
}

// Controllers
func (server *Server) createShop(ctx *gin.Context) {
	var req createShopRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.CreateShopParams{
		ShopName: req.ShopName,
		ShopType: req.ShopType,
		Address:  req.Address,
		UserID:   req.UserID,
	}

	user, err := server.store.CreateShop(ctx, args)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation", "foreign_key_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			default:
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
		}
	}
	ctx.JSON(http.StatusCreated, user)
}

func (server *Server) updateShopById(ctx *gin.Context) {
	var req UpdateShopDetailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.UpdateShopDetailParams{
		ShopName: req.ShopName,
		ShopType: req.ShopType,
		Address:  req.Address,
		ID:       req.ID,
	}

	updatedshop, err := server.store.UpdateShopDetail(ctx, args)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, updatedshop)
}

func (server *Server) getShopById(ctx *gin.Context) {
	var req getShopByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	shop, err := server.store.GetShopById(ctx, int32(req.ID))

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, shop)
}

func (server *Server) countShopById(ctx *gin.Context) {
	var req getShopByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	shop, err := server.store.CountShops(ctx, int32(req.ID))

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	resp := StatResponse{
		Count: shop,
	}
	ctx.JSON(http.StatusOK, resp)
}

func (server *Server) deleteShopById(ctx *gin.Context) {
	var req deleteShopByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteShop(ctx, req.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, successResponse(http.StatusOK, "Shop Deleted Successfully"))
}

func (server *Server) listMyShops(ctx *gin.Context) {
	var req listMyShopsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.ListMyShopsParams{
		UserID: req.UserID,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	shops, err := server.store.ListMyShops(ctx, args)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, shops)
}
