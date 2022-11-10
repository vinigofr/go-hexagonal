package application_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/vinigofr/go-hexagonal/application"
	mock_application "github.com/vinigofr/go-hexagonal/application/mocks"
)

func TestProductService_Get(t *testing.T) {
	// Instanciamos o ctrl, é o controller do gomock para podermor usar no
	// arquivo de mock gerado via mockgen
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Instanciamos um product falso
	product := mock_application.NewMockProductInterface(ctrl)

	// Instanciamos uma interface persistence falsa
	persistense := mock_application.NewMockProductPersistenceInterface(ctrl)

	// Dizemos o que queremos que seja retornado a cada vez que o método Get()
	// De persistence for chamado
	persistense.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistense,
	}

	result, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)

}

func TestProductService_Save(t *testing.T) {
	// Instanciamos o ctrl, é o controller do gomock para podermor usar no
	// arquivo de mock gerado via mockgen
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Instanciamos um product falso
	product := mock_application.NewMockProductInterface(ctrl)

	// Instanciamos uma interface persistence falsa
	persistense := mock_application.NewMockProductPersistenceInterface(ctrl)

	// Dizemos o que queremos que seja retornado a cada vez que o método Get()
	// De persistence for chamado
	persistense.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistense,
	}

	// Criamos um produto
	result, err := service.Create("Produto 1", 1)

	require.Nil(t, err)
	require.Equal(t, product, result)

}

func TestProductService_Enable(t *testing.T) {
	// Instanciamos o ctrl, é o controller do gomock para podermor usar no
	// arquivo de mock gerado via mockgen
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Instanciamos um product falso
	product := mock_application.NewMockProductInterface(ctrl)

	// Instanciamos uma interface persistence falsa
	persistense := mock_application.NewMockProductPersistenceInterface(ctrl)

	// Dizemos o que queremos que seja retornado a cada vez que o método Get()
	// De persistence for chamado
	product.EXPECT().Enable().Return(nil).AnyTimes()
	persistense.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistense,
	}

	result, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}
