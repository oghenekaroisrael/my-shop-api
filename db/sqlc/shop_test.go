package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/oghenekaroisrael/myshopapi/utils"
	"github.com/stretchr/testify/require"
)

func createRandomShop(t *testing.T, userId int32) Shop {
	arg := CreateShopParams{
		ShopName: utils.RandomName(),
		ShopType: utils.RandomName(),
		Address:  utils.RandomName(),
		UserID:   userId,
	}
	shop, err := testQueries.CreateShop(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, shop)
	require.Equal(t, arg.ShopName, shop.ShopName)
	require.Equal(t, arg.ShopType, shop.ShopType)
	require.Equal(t, arg.Address, shop.Address)
	require.Equal(t, arg.UserID, shop.UserID)
	require.NotZero(t, shop.ID)
	require.NotZero(t, shop.CreatedAt)
	return shop
}

func TestCreateShop(t *testing.T) {
	user := createRandomUser(t)
	createRandomShop(t, user.ID)
}

func TestGetShopById(t *testing.T) {
	user := createRandomUser(t)
	shop1 := createRandomShop(t, user.ID)
	shop2, err := testQueries.GetShopById(context.Background(), shop1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, shop2)
	require.Equal(t, shop1.ID, shop2.ID)
	require.Equal(t, shop1.ShopName, shop2.ShopName)
	require.Equal(t, shop1.ShopType, shop2.ShopType)
	require.Equal(t, shop1.Address, shop2.Address)
	require.Equal(t, shop1.UserID, shop2.UserID)
	require.WithinDuration(t, shop1.CreatedAt, shop2.CreatedAt, time.Second)
}

func TestDeleteShop(t *testing.T) {
	user := createRandomUser(t)
	shop1 := createRandomShop(t, user.ID)

	err := testQueries.DeleteShop(context.Background(), shop1.ID)
	require.NoError(t, err)

	shops, err := testQueries.GetShopById(context.Background(), shop1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, shops)
}

func TestUpdateShopDetail(t *testing.T) {

	user := createRandomUser(t)
	shop1 := createRandomShop(t, user.ID)
	args := UpdateShopDetailParams{
		ID:       shop1.ID,
		ShopName: utils.RandomName(),
		ShopType: utils.RandomName(),
		Address:  utils.RandomName(),
	}
	shop2, err := testQueries.UpdateShopDetail(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, shop2)
	require.Equal(t, args.ID, shop2.ID)
	require.Equal(t, args.ShopName, shop2.ShopName)
	require.Equal(t, args.ShopType, shop2.ShopType)
	require.Equal(t, args.Address, shop2.Address)
	require.WithinDuration(t, shop1.CreatedAt, shop2.CreatedAt, time.Second)
}

func TestListMyShops(t *testing.T) {
	user := createRandomUser(t)
	for i := 0; i < 10; i++ {
		createRandomShop(t, user.ID)
	}

	args := ListMyShopsParams{
		UserID: user.ID,
		Limit:  5,
		Offset: 5,
	}

	shops, err := testQueries.ListMyShops(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, shops, 5)
	for _, shop := range shops {
		require.NotEmpty(t, shop)
	}
}
