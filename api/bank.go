package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/oghenekaroisrael/myshopapi/db/sqlc"
)

// Request Structs
type createBankRequest struct {
	BankName      string `json:"bank_name" binding:"required"`
	AccountNumber string `json:"account_number" binding:"required,min=10,max=10"`
	Icon          string `json:"icon" binding:"required"`
	ShopID        int32  `json:"shop_id" binding:"required"`
}

type getBankByIDRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type deleteBankByIDRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type listMyBanksRequest struct {
	ShopID   int32 `form:"shop_id" binding:"required"`
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}

type UpdateBankDetailRequest struct {
	ID            int32  `json:"id" binding:"required"`
	BankName      string `json:"bank_name" binding:"required"`
	AccountNumber string `json:"account_number" binding:"required,min=10,max=10"`
	Icon          string `json:"icon" binding:"required"`
}

// Controllers
func (server *Server) createBank(ctx *gin.Context) {
	var req createBankRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.CreateBankParams{
		BankName:      req.BankName,
		AccountNumber: req.AccountNumber,
		Icon:          req.Icon,
		ShopID:        req.ShopID,
	}

	bank, err := server.store.CreateBank(ctx, args)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, bank)
}

func (server *Server) updateBankById(ctx *gin.Context) {
	var req UpdateBankDetailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.UpdateBankParams{
		BankName:      req.BankName,
		AccountNumber: req.AccountNumber,
		Icon:          req.Icon,
		ID:            req.ID,
	}

	updatedbank, err := server.store.UpdateBank(ctx, args)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, updatedbank)
}

func (server *Server) getBankById(ctx *gin.Context) {
	var req getBankByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	bank, err := server.store.GetBankById(ctx, int32(req.ID))

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, bank)
}

func (server *Server) deleteBankById(ctx *gin.Context) {
	var req deleteBankByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteBank(ctx, req.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, successResponse(http.StatusOK, "Bank Deleted Successfully"))
}

func (server *Server) listMyBanks(ctx *gin.Context) {
	var req listMyBanksRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.ListMyBanksParams{
		ShopID: req.ShopID,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	banks, err := server.store.ListMyBanks(ctx, args)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, banks)
}
