/*
Para testar funções internas
package application
*/

// Para testar funções externas
package application_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/vinigofr/go-hexagonal/application"
)

func generateNewProduct(p application.Product) application.Product {
	return p
}

func TestProduct_Enable(t *testing.T) {
	product := application.Product{
		Name:   "Carro",
		Price:  10,
		Status: application.DISABLED,
	}

	errorMessage := "the price must be greather than zero to enable the product"
	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()

	require.Equal(t, errorMessage, err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{
		Name:   "Carro",
		Price:  0,
		Status: application.ENABLED,
	}

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 1
	errorMessage := "the price must be equal to zero to disable the product"
	err = product.Disable()
	require.Equal(t, errorMessage, err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product1 := generateNewProduct(application.Product{
		Name:   "Carro",
		Price:  1.0,
		Status: application.DISABLED,
		ID:     uuid.NewV4().String(),
	})

	product2 := generateNewProduct(application.Product{
		Name:   "Carro",
		Price:  0,
		Status: application.DISABLED,
		ID:     uuid.NewV4().String(),
	})

	// Valid status testing
	_, err := product1.IsValid()
	require.Nil(t, err)

	// Invalid status testing
	product1.Status = "INVALID_STATUS"
	_, err = product1.IsValid()
	require.Equal(t, "status must be enabled or disabled", err.Error())

	// Price with zero value
	_, err = product2.IsValid()
	require.Equal(t, "the price must be greather or equal zero", err.Error())

	// Invalid struct
	product1 = application.Product{}
	boolean, err := product1.IsValid()
	require.Equal(t, boolean, false)
	require.Error(t, err)

}
