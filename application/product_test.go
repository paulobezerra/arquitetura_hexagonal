package application_test

import (
	"testing"

	"github.com/bezerrapaulo/go-hexagonal/application"
	"github.com/stretchr/testify/require"

	uuid "github.com/satori/go.uuid"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Test"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()

	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Test"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()

	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be zero to disable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Test"
	product.Status = application.DISABLED
	product.Price = 10.0

	_, err := product.IsValid()

	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater than zero", err.Error())

	product.Price = 10
	product.Status = "invalid"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.ID = "abc"
	_, err = product.IsValid()
	require.Equal(t, "id: abc does not validate as uuidv4", err.Error())
}

func TestProduct_Getters(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Test"
	product.Status = application.DISABLED
	product.Price = 10.0

	require.Equal(t, product.ID, product.GetID())
	require.Equal(t, product.Name, product.GetName())
	require.Equal(t, product.Status, product.GetStatus())
	require.Equal(t, product.Price, product.GetPrice())
}

func TestProduct_NewProduct(t *testing.T) {
	product := application.NewProduct()
	require.NotEmpty(t, product.ID)
	require.Equal(t, application.DISABLED, product.Status)
}
