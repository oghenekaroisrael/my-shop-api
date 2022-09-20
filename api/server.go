package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/oghenekaroisrael/myshopapi/db/sqlc"
	// "github.com/oghenekaroisrael/myshopapi/token"
)

type Server struct {
	store db.Store
	// tokenMaker token.Maker
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	// tokenMaker, err := token.NewPasetoMaker()
	// if err != nil {
	// 	return nil, fmt.Errorf("cannot create token %v", err)
	// }
	server := &Server{
		store: store,
		// tokenMaker: tokenMaker,
	}
	router := gin.Default()
	router.POST("/users/signup", server.createUser)
	router.POST("users/login", server.loginUser)
	router.GET("users/:id", server.getUserById)
	router.GET("users", server.listUsers)

	router.POST("/users/shops", server.createShop)
	router.GET("users/shops/:id", server.getShopById)
	router.GET("users/shops/count/:id", server.countShopById)
	router.DELETE("/users/shops/:id", server.deleteShopById)
	router.PATCH("/users/shops", server.updateShopById)
	router.GET("users/shops", server.listMyShops)

	router.POST("/users/shops/banks", server.createBank)
	router.GET("users/shops/banks/:id", server.getBankById)
	router.DELETE("/users/shops/banks/:id", server.deleteBankById)
	router.PATCH("/users/shops/banks", server.updateBankById)
	router.GET("users/shops/banks", server.listMyBanks)

	router.POST("/users/shops/inventory", server.createInventory)
	router.GET("users/shops/inventory/:id", server.getItemById)
	router.DELETE("/users/shops/inventory/:id", server.deleteItemById)
	router.PATCH("/users/shops/inventory", server.updateItemById)
	router.PATCH("/users/shops/inventory/quantity", server.updateItemQuantityById)
	router.GET("users/shops/inventory", server.listMyInventoryByShop)

	router.POST("/users/shops/sales", server.newSale)
	router.GET("users/shops/sales/:id", server.getSaleById)
	router.DELETE("/users/shops/sales/:id", server.deleteSaleById)
	router.GET("users/shops/sales", server.listShopSales)
	router.GET("users/shops/sales/user", server.listSalesByUser)

	router.POST("/users/shops/expenses", server.createExpenses)
	router.GET("users/shops/expenses/:id", server.getExpensesById)
	router.DELETE("/users/shops/expenses/:id", server.deleteExpensesById)
	router.GET("users/shops/expenses", server.listMyShopExpenses)
	router.GET("users/shops/expenses/dates", server.getExpensesBetweenDates)

	router.POST("/users/shops/payouts", server.createPayout)
	router.GET("users/shops/payouts/:id", server.getPayoutById)
	router.DELETE("/users/shops/payouts/:id", server.deletePayout)
	router.GET("users/shops/payouts", server.listMyShopPayouts)
	router.GET("users/shops/payouts/dates", server.getPayoutBetweenDates)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func successResponse(code int, res string) gin.H {
	return gin.H{
		"payload": res,
		"code":    code,
	}
}
