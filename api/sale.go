package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/oghenekaroisrael/myshopapi/db/sqlc"
)

// Request Structs
type createSaleRequest struct {
	ItemID             int32  `json:"item_id" binding:"required"`
	Quantity           int32  `json:"quantity" binding:"required,gt=0"`
	SellingPriceActual int32  `json:"selling_price_actual" binding:"required,gt=0"`
	PaymentType        string `json:"payment_type" binding:"required"`
	ShopID             int32  `json:"shop_id" binding:"required"`
}

type getSaleByIDRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type deleteSaleByIDRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type listMyShopSaleRequest struct {
	ShopID   int32 `form:"shop_id" binding:"required"`
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}

type listSaleByUserRequest struct {
	UserID   int32 `form:"user_id" binding:"required"`
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}

// Controllers
func (server *Server) newSale(ctx *gin.Context) {
	var req createSaleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.CreateSaleParams{
		ItemID:             req.ItemID,
		Quantity:           req.Quantity,
		SellingPriceActual: req.SellingPriceActual,
		PaymentType:        req.PaymentType,
		ShopID:             req.ShopID,
	}

	sale, err := server.store.CreateSale(ctx, args)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, sale)
}

func (server *Server) getSaleById(ctx *gin.Context) {
	var req getSaleByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	sale, err := server.store.GetSaleById(ctx, int32(req.ID))

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, sale)
}

func (server *Server) deleteSaleById(ctx *gin.Context) {
	var req deleteSaleByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteSale(ctx, req.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, successResponse(http.StatusOK, "Sale Deleted Successfully"))
}

func (server *Server) listShopSales(ctx *gin.Context) {
	var req listMyShopSaleRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.ListMySalesParams{
		ShopID: req.ShopID,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	sales, err := server.store.ListMySales(ctx, args)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, sales)
}

func (server *Server) listSalesByUser(ctx *gin.Context) {
	var req listSaleByUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.ListSalesByUserParams{
		ID:     req.UserID,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	sales, err := server.store.ListSalesByUser(ctx, args)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, sales)
}
