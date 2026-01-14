package product

import (
	"errors"
	"fmt"
	"testing"

	"github.com/imbafff/product-warehouse-api/internal/entity"
)

// Mock Repository для тестирования
type MockRepository struct {
	products map[int64]*entity.Product
	nextID   int64
}

func NewMockRepository() *MockRepository {
	return &MockRepository{
		products: make(map[int64]*entity.Product),
		nextID:   1,
	}
}

func (m *MockRepository) Create(p *entity.Product) (int64, error) {
	id := m.nextID
	p.ID = id
	m.products[id] = p
	m.nextID++
	return id, nil
}

func (m *MockRepository) GetByID(id int64) (*entity.Product, error) {
	if p, exists := m.products[id]; exists {
		return p, nil
	}
	return nil, errors.New("product not found")
}

func (m *MockRepository) Update(id int64, p *entity.Product) error {
	if _, exists := m.products[id]; !exists {
		return errors.New("product not found")
	}
	p.ID = id
	m.products[id] = p
	return nil
}

func (m *MockRepository) Delete(id int64) error {
	if _, exists := m.products[id]; !exists {
		return errors.New("product not found")
	}
	delete(m.products, id)
	return nil
}

func (m *MockRepository) GetAll() ([]*entity.Product, error) {
	products := make([]*entity.Product, 0, len(m.products))
	for _, p := range m.products {
		products = append(products, p)
	}
	return products, nil
}

// Тесты для Create
func TestCreate_Success(t *testing.T) {
	repo := NewMockRepository()
	service := New(repo)

	product := &entity.Product{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       10.99,
		Quantity:    5,
	}

	id, err := service.Create(product)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if id != 1 {
		t.Errorf("Expected id 1, got %d", id)
	}

	// Проверяем, что продукт был создан
	created, err := service.GetByID(id)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if created.Name != "Test Product" {
		t.Errorf("Expected name 'Test Product', got %s", created.Name)
	}
}

func TestCreate_EmptyName(t *testing.T) {
	repo := NewMockRepository()
	service := New(repo)

	product := &entity.Product{
		Name:     "",
		Price:    10.99,
		Quantity: 5,
	}

	_, err := service.Create(product)
	if err == nil {
		t.Error("Expected error for empty name, got nil")
	}

	if err.Error() != "name is required" {
		t.Errorf("Expected 'name is required', got %v", err)
	}
}

func TestCreate_InvalidPrice(t *testing.T) {
	repo := NewMockRepository()
	service := New(repo)

	testCases := []float64{0, -10.5, -0.01}

	for _, price := range testCases {
		product := &entity.Product{
			Name:     "Test",
			Price:    price,
			Quantity: 5,
		}

		_, err := service.Create(product)
		if err == nil {
			t.Errorf("Expected error for price %.2f, got nil", price)
		}

		if err.Error() != "price must be greater than zero" {
			t.Errorf("Expected 'price must be greater than zero', got %v", err)
		}
	}
}

func TestCreate_NegativeQuantity(t *testing.T) {
	repo := NewMockRepository()
	service := New(repo)

	product := &entity.Product{
		Name:     "Test",
		Price:    10.99,
		Quantity: -5,
	}

	_, err := service.Create(product)
	if err == nil {
		t.Error("Expected error for negative quantity, got nil")
	}

	if err.Error() != "quantity must be non-negative" {
		t.Errorf("Expected 'quantity must be non-negative', got %v", err)
	}
}

// Тесты для GetByID
func TestGetByID_Success(t *testing.T) {
	repo := NewMockRepository()
	service := New(repo)

	product := &entity.Product{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       10.99,
		Quantity:    5,
	}

	id, _ := service.Create(product)

	retrieved, err := service.GetByID(id)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if retrieved.Name != "Test Product" {
		t.Errorf("Expected name 'Test Product', got %s", retrieved.Name)
	}

	if retrieved.Price != 10.99 {
		t.Errorf("Expected price 10.99, got %f", retrieved.Price)
	}
}

func TestGetByID_NotFound(t *testing.T) {
	repo := NewMockRepository()
	service := New(repo)

	_, err := service.GetByID(999)
	if err == nil {
		t.Error("Expected error for non-existent product, got nil")
	}

	if err.Error() != "product not found" {
		t.Errorf("Expected 'product not found', got %v", err)
	}
}

func TestGetByID_InvalidID(t *testing.T) {
	repo := NewMockRepository()
	service := New(repo)

	testCases := []int64{0, -1, -999}

	for _, id := range testCases {
		_, err := service.GetByID(id)
		if err == nil {
			t.Errorf("Expected error for id %d, got nil", id)
		}

		if err.Error() != "invalid id" {
			t.Errorf("Expected 'invalid id', got %v", err)
		}
	}
}

// Тесты для Update
func TestUpdate_Success(t *testing.T) {
	repo := NewMockRepository()
	service := New(repo)

	product := &entity.Product{
		Name:        "Original",
		Description: "Original Description",
		Price:       10.99,
		Quantity:    5,
	}

	id, _ := service.Create(product)

	updatedProduct := &entity.Product{
		Name:        "Updated",
		Description: "Updated Description",
		Price:       20.99,
		Quantity:    10,
	}

	err := service.Update(id, updatedProduct)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	retrieved, _ := service.GetByID(id)
	if retrieved.Name != "Updated" {
		t.Errorf("Expected name 'Updated', got %s", retrieved.Name)
	}

	if retrieved.Price != 20.99 {
		t.Errorf("Expected price 20.99, got %f", retrieved.Price)
	}
}

func TestUpdate_NotFound(t *testing.T) {
	repo := NewMockRepository()
	service := New(repo)

	updatedProduct := &entity.Product{
		Name:     "Updated",
		Price:    20.99,
		Quantity: 10,
	}

	err := service.Update(999, updatedProduct)
	if err == nil {
		t.Error("Expected error for non-existent product, got nil")
	}

	if err.Error() != "product not found" {
		t.Errorf("Expected 'product not found', got %v", err)
	}
}

func TestUpdate_InvalidData(t *testing.T) {
	repo := NewMockRepository()
	service := New(repo)

	product := &entity.Product{
		Name:     "Test",
		Price:    10.99,
		Quantity: 5,
	}

	id, _ := service.Create(product)

	testCases := []struct {
		name    string
		product *entity.Product
		errMsg  string
	}{
		{
			name: "empty name",
			product: &entity.Product{
				Name:     "",
				Price:    10.99,
				Quantity: 5,
			},
			errMsg: "name is required",
		},
		{
			name: "invalid price",
			product: &entity.Product{
				Name:     "Test",
				Price:    -5.0,
				Quantity: 5,
			},
			errMsg: "price must be greater than zero",
		},
		{
			name: "negative quantity",
			product: &entity.Product{
				Name:     "Test",
				Price:    10.99,
				Quantity: -5,
			},
			errMsg: "quantity must be non-negative",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := service.Update(id, tc.product)
			if err == nil {
				t.Error("Expected error, got nil")
			}

			if err.Error() != tc.errMsg {
				t.Errorf("Expected '%s', got %v", tc.errMsg, err)
			}
		})
	}
}

// Тесты для Delete
func TestDelete_Success(t *testing.T) {
	repo := NewMockRepository()
	service := New(repo)

	product := &entity.Product{
		Name:     "Test",
		Price:    10.99,
		Quantity: 5,
	}

	id, _ := service.Create(product)

	err := service.Delete(id)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	_, err = service.GetByID(id)
	if err == nil {
		t.Error("Expected error when getting deleted product, got nil")
	}
}

func TestDelete_NotFound(t *testing.T) {
	repo := NewMockRepository()
	service := New(repo)

	err := service.Delete(999)
	if err == nil {
		t.Error("Expected error for non-existent product, got nil")
	}

	if err.Error() != "product not found" {
		t.Errorf("Expected 'product not found', got %v", err)
	}
}

func TestDelete_InvalidID(t *testing.T) {
	repo := NewMockRepository()
	service := New(repo)

	testCases := []int64{0, -1, -999}

	for _, id := range testCases {
		err := service.Delete(id)
		if err == nil {
			t.Errorf("Expected error for id %d, got nil", id)
		}

		if err.Error() != "invalid id" {
			t.Errorf("Expected 'invalid id', got %v", err)
		}
	}
}

// Тесты для GetAll
func TestGetAll_Empty(t *testing.T) {
	repo := NewMockRepository()
	service := New(repo)

	products, err := service.GetAll()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(products) != 0 {
		t.Errorf("Expected 0 products, got %d", len(products))
	}
}

func TestGetAll_MultipleProducts(t *testing.T) {
	repo := NewMockRepository()
	service := New(repo)

	// Создаем несколько продуктов
	for i := 1; i <= 3; i++ {
		product := &entity.Product{
			Name:     fmt.Sprintf("Product %d", i),
			Price:    float64(10*i) + 0.99,
			Quantity: i * 5,
		}
		service.Create(product)
	}

	products, err := service.GetAll()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(products) != 3 {
		t.Errorf("Expected 3 products, got %d", len(products))
	}
}
