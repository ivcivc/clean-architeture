package usecase

import (
	"testing"

	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/stretchr/testify/assert"
)

type MockOrderRepository struct {
	orders []entity.Order
}

func (m *MockOrderRepository) Save(order *entity.Order) error {
	return nil
}

func (m *MockOrderRepository) List() ([]entity.Order, error) {
	return m.orders, nil
}

func TestListOrdersUseCase_Execute(t *testing.T) {
	mockRepo := &MockOrderRepository{
		orders: []entity.Order{
			{ID: "1", Price: 100.0, Tax: 10.0, FinalPrice: 110.0},
			{ID: "2", Price: 200.0, Tax: 20.0, FinalPrice: 220.0},
		},
	}
	usecase := NewListOrdersUseCase(mockRepo)

	result, err := usecase.Execute()
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "1", result[0].ID)
	assert.Equal(t, 100.0, result[0].Price)
	assert.Equal(t, 10.0, result[0].Tax)
	assert.Equal(t, 110.0, result[0].FinalPrice)
	assert.Equal(t, "2", result[1].ID)
	assert.Equal(t, 200.0, result[1].Price)
	assert.Equal(t, 20.0, result[1].Tax)
	assert.Equal(t, 220.0, result[1].FinalPrice)
}
