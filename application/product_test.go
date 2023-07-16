package application_test

import (
	"ports-and-adapters/application"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Product Test"
	product.Price = 10
	product.Status = application.DISABLED

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Product Test"
	product.Price = 0
	product.Status = application.ENABLED

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.Name = "Product Test"
	product.Price = 10
	product.Status = application.DISABLED
	product.ID = uuid.NewV4().String()

	isValid, err := product.IsValid()
	require.Nil(t, err)
	require.True(t, isValid)

	product.Status = "INVALID_STATUS"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.DISABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal zero", err.Error())
}

func TestProduct_GetId(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	require.Equal(t, product.ID, product.GetID())
}

func TestProduct_GetName(t *testing.T) {
	product := application.Product{}
	product.Name = "Product Test"

	require.Equal(t, product.Name, product.GetName())
}

func TestProduct_GetStatus(t *testing.T) {
	product := application.Product{}
	product.Status = application.DISABLED

	require.Equal(t, product.Status, product.GetStatus())
}

func TestProduct_GetPrice(t *testing.T) {
	product := application.Product{}
	product.Price = 10

	require.Equal(t, product.Price, product.GetPrice())
}
