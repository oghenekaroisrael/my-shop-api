package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/oghenekaroisrael/myshopapi/db/sqlc"
)

// Request Structs
type createInventoryRequest struct {
	ItemName             string `json:"item_name" binding:"required"`
	Quantity             int32  `json:"quantity" binding:"required,gt=0"`
	CostPrice            int32  `json:"cost_price" binding:"required,gt=0"`
	SellingPriceStandard int32  `json:"selling_price_standard" binding:"required,gt=0"`
	Status               string `json:"status" binding:"required"`
	ShopID               int32  `json:"shop_id" binding:"required"`
}

type getInventoryByIDRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type deleteInventoryByIDRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type listMyInventoryByShopRequest struct {
	ShopID   int32 `form:"shop_id" binding:"required"`
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}

type UpdateInventoryDetailRequest struct {
	ID                   int32  `json:"id" binding:"required"`
	ItemName             string `json:"item_name" binding:"required"`
	CostPrice            int32  `json:"cost_price" binding:"required,gt=0"`
	SellingPriceStandard int32  `json:"selling_price_standard" binding:"required,gt=0"`
	Status               string `json:"status" binding:"required"`
}

type UpdateInventoryQuantityRequest struct {
	ID       int32 `json:"id" binding:"required"`
	Quantity int32 `json:"quantity" binding:"required,gt=0"`
}

// Controllers
func (server *Server) createInventory(ctx *gin.Context) {
	var req createInventoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.CreateItemParams{
		ItemName:             req.ItemName,
		Quantity:             req.Quantity,
		SellingPriceStandard: req.SellingPriceStandard,
		CostPrice:            req.CostPrice,
		Status:               req.Status,
		ShopID:               req.ShopID,
	}

	item, err := server.store.CreateItem(ctx, args)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, item)
}

func (server *Server) updateItemById(ctx *gin.Context) {
	var req UpdateInventoryDetailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.UpdateItemDetailParams{
		ItemName:             req.ItemName,
		CostPrice:            req.CostPrice,
		SellingPriceStandard: req.SellingPriceStandard,
		ID:                   req.ID,
		Status:               req.Status,
	}

	updateditem, err := server.store.UpdateItemDetail(ctx, args)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, updateditem)
}

func (server *Server) updateItemQuantityById(ctx *gin.Context) {
	var req UpdateInventoryQuantityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.UpdateItemQuantityParams{
		ID:       req.ID,
		Quantity: req.Quantity,
	}

	updateditem, err := server.store.UpdateItemQuantity(ctx, args)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, updateditem)
}

func (server *Server) getItemById(ctx *gin.Context) {
	var req getInventoryByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	item, err := server.store.GetItemById(ctx, int32(req.ID))

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, item)
}

func (server *Server) deleteItemById(ctx *gin.Context) {
	var req deleteInventoryByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteItem(ctx, req.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, successResponse(http.StatusOK, "Item Deleted From Inventory Successfully"))
}

func (server *Server) listMyInventoryByShop(ctx *gin.Context) {
	var req listMyInventoryByShopRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.ListMyInventoriesParams{
		ShopID: req.ShopID,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	items, err := server.store.ListMyInventories(ctx, args)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, items)
}
