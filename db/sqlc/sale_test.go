package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/oghenekaroisrael/myshopapi/utils"
	"github.com/stretchr/testify/require"
)

func createRandomSale(t *testing.T) Sale {
	arg := CreateSaleParams{
		ItemID:             utils.RandomInt(1, 5),
		Quantity:           utils.RandomInt(1, 5),
		SellingPriceActual: utils.RandomInt(0, 1000),
		PaymentType:        "Cash",
		ShopID:             utils.RandomInt(1, 5),
	}
	sale, err := testQueries.CreateSale(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, sale)
	require.Equal(t, arg.ItemID, sale.ItemID)
	require.Equal(t, arg.Quantity, sale.Quantity)
	require.Equal(t, arg.SellingPriceActual, sale.SellingPriceActual)
	require.Equal(t, arg.PaymentType, sale.PaymentType)
	require.NotZero(t, sale.ID)
	require.NotZero(t, sale.CreatedAt)
	return sale
}

func TestCreateSale(t *testing.T) {
	createRandomSale(t)
}

func TestGetSaleById(t *testing.T) {
	sale1 := createRandomSale(t)
	sale2, err := testQueries.GetSaleById(context.Background(), sale1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, sale2)
	require.Equal(t, sale1.ID, sale2.ID)
	require.Equal(t, sale1.ItemID, sale2.ItemID)
	require.Equal(t, sale1.Quantity, sale2.Quantity)
	require.Equal(t, sale1.SellingPriceActual, sale2.SellingPriceActual)
	require.Equal(t, sale1.PaymentType, sale2.PaymentType)
	require.WithinDuration(t, sale1.CreatedAt, sale2.CreatedAt, time.Second)
}

// func TestGetSearchSales(t *testing.T) {
// 	sale1 := createRandomSale(t)
// 	args := SearchSalesParams{
// 		ShopID:   sale1.ShopID,
// 		SaleName: sale1.SaleName,
// 		Limit:    5,
// 		Offset:   5,
// 	}
// 	sales, err := testQueries.SearchSales(context.Background(), args)
// 	require.NoError(t, err)
// 	require.Len(t, sales, 5)
// 	for _, sale := range sales {
// 		require.NotEmpty(t, sale)
// 		require.Equal(t, sale1.ID, sale.ID)
// 		require.Equal(t, sale1.SaleName, sale.SaleName)
// 		require.Equal(t, sale1.SaleType, sale.SaleType)
// 		require.Equal(t, sale1.Address, sale.Address)
// 		require.Equal(t, sale1.ShopID, sale.ShopID)
// 		require.WithinDuration(t, sale1.CreatedAt, sale.CreatedAt, time.Second)
// 	}
// }

func TestDeleteSale(t *testing.T) {
	sale1 := createRandomSale(t)

	err := testQueries.DeleteSale(context.Background(), sale1.ID)
	require.NoError(t, err)

	sales, err := testQueries.GetSaleById(context.Background(), sale1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, sales)
}

// func TestUpdateSaleDetail(t *testing.T) {
// 	sale1 := createRandomSale(t)
// 	args := UpdateSaleDetailParams{
// 		ID:       sale1.ID,
// 		SaleName: utils.RandomName(),
// 		SaleType: utils.RandomName(),
// 		Address:  utils.RandomName(),
// 	}
// 	sale2, err := testQueries.UpdateSaleDetail(context.Background(), args)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, sale2)
// 	require.Equal(t, args.ID, sale2.ID)
// 	require.Equal(t, args.SaleName, sale2.SaleName)
// 	require.Equal(t, args.SaleType, sale2.SaleType)
// 	require.Equal(t, args.Address, sale2.Address)
// 	require.WithinDuration(t, sale1.CreatedAt, sale2.CreatedAt, time.Second)
// }

func TestListMySales(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomSale(t)
	}

	args := ListMySalesParams{
		ShopID: 1,
		Limit:  5,
		Offset: 5,
	}

	sales, err := testQueries.ListMySales(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, sales, 5)
	for _, sale := range sales {
		require.NotEmpty(t, sale)
	}
}
