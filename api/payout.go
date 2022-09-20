package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/oghenekaroisrael/myshopapi/db/sqlc"
)

// Request Structs
type createPayoutRequest struct {
	Particular  string `json:"particular" binding:"required"`
	Amount      int32  `json:"amount" binding:"required,gt=0"`
	Recipient   string `json:"recipient" binding:"required"`
	PaymentType string `json:"payment_type" binding:"required"`
	ShopID      int32  `json:"shop_id" binding:"required"`
}

type getPayoutRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type deletePayoutRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type searchPayoutBetweenDatesRequest struct {
	ShopID   int32     `json:"shop_id" binding:"required"`
	FromDate time.Time `json:"from_date" binding:"required"`
	ToDate   time.Time `json:"to_date" binding:"required"`
	Limit    int32     `json:"limit" binding:"required"`
	Offset   int32     `json:"offset" binding:"required"`
}

type listMyPayoutsRequest struct {
	ShopID   int32 `form:"shop_id" binding:"required"`
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}

// Controllers
func (server *Server) createPayout(ctx *gin.Context) {
	var req createPayoutRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.CreatePayoutParams{
		Particular:  req.Particular,
		Recipient:   req.Recipient,
		Amount:      req.Amount,
		PaymentType: req.PaymentType,
		ShopID:      req.ShopID,
	}

	payout, err := server.store.CreatePayout(ctx, args)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, payout)
}

func (server *Server) getPayoutById(ctx *gin.Context) {
	var req getPayoutRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payout, err := server.store.GetPayoutById(ctx, int32(req.ID))

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, payout)
}

func (server *Server) getPayoutBetweenDates(ctx *gin.Context) {
	var req searchPayoutBetweenDatesRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	args := db.SearchBetweenDatesParams{
		CreatedAt:   req.FromDate,
		CreatedAt_2: req.ToDate,
		Offset:      req.Offset,
		ShopID:      req.ShopID,
		Limit:       req.Limit,
	}

	payouts, err := server.store.SearchBetweenDates(ctx, args)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, payouts)
}

func (server *Server) deletePayout(ctx *gin.Context) {
	var req deletePayoutRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeletePayout(ctx, req.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, successResponse(http.StatusOK, "Payout Deleted Successfully"))
}

func (server *Server) listMyShopPayouts(ctx *gin.Context) {
	var req listMyPayoutsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.ListMyPayoutsParams{
		ShopID: req.ShopID,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	payouts, err := server.store.ListMyPayouts(ctx, args)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, payouts)
}
