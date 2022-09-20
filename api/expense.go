package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/oghenekaroisrael/myshopapi/db/sqlc"
)

// Request Structs
type createExpenseRequest struct {
	Particular  string `json:"particular" binding:"required"`
	Amount      int32  `json:"amount" binding:"required,gt=0"`
	Recipient   string `json:"recipient" binding:"required"`
	PaymentType string `json:"payment_type" binding:"required"`
	ShopID      int32  `json:"shop_id" binding:"required"`
}

type getExpenseByIDRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type deleteExpenseByIDRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type searchExpensesBetweenDatesRequest struct {
	ShopID   int32     `json:"shop_id" binding:"required"`
	FromDate time.Time `json:"from_date" binding:"required"`
	ToDate   time.Time `json:"to_date" binding:"required"`
	Limit    int32     `json:"limit" binding:"required"`
	Offset   int32     `json:"offset" binding:"required"`
}

type listMyExpensesRequest struct {
	ShopID   int32 `form:"shop_id" binding:"required"`
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}

// Controllers
func (server *Server) createExpenses(ctx *gin.Context) {
	var req createExpenseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.CreateExpensesParams{
		Particular:  req.Particular,
		Recipient:   req.Recipient,
		Amount:      req.Amount,
		PaymentType: req.PaymentType,
		ShopID:      req.ShopID,
	}

	exp, err := server.store.CreateExpenses(ctx, args)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, exp)
}

func (server *Server) getExpensesById(ctx *gin.Context) {
	var req getExpenseByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	exp, err := server.store.GetExpenseById(ctx, int32(req.ID))

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, exp)
}

func (server *Server) getExpensesBetweenDates(ctx *gin.Context) {
	var req searchExpensesBetweenDatesRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	args := db.SearchExpensesBetweenDatesParams{
		CreatedAt:   req.FromDate,
		CreatedAt_2: req.ToDate,
		Offset:      req.Offset,
		ShopID:      req.ShopID,
		Limit:       req.Limit,
	}

	exps, err := server.store.SearchExpensesBetweenDates(ctx, args)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, exps)
}

func (server *Server) deleteExpensesById(ctx *gin.Context) {
	var req deleteExpenseByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteExpense(ctx, req.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, successResponse(http.StatusOK, "Expense Deleted Successfully"))
}

func (server *Server) listMyShopExpenses(ctx *gin.Context) {
	var req listMyExpensesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.ListMyExpensesParams{
		ShopID: req.ShopID,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	exps, err := server.store.ListMyExpenses(ctx, args)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, exps)
}
