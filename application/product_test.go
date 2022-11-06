/*
Para testar funções internas
package application
*/

// Para testar funções externas
package application_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vinigofr/go-hexagonal/application"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{
		Name:   "Carro",
		Price:  10,
		Status: application.DISABLED,
	}

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greather than zero to enable the product", err.Error())
}
