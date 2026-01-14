package product

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/imbafff/product-warehouse-api/internal/entity"
	"github.com/imbafff/product-warehouse-api/internal/infrastructure/config"
	"github.com/imbafff/product-warehouse-api/internal/infrastructure/db"
	productRepo "github.com/imbafff/product-warehouse-api/internal/repository/product"
)

// Интеграционные тесты с реальной БД (если БД доступна)
func getTestDB(t *testing.T) *sql.DB {
	cfg := config.Load()
	database, err := db.NewPostgresDB(cfg)
	if err != nil {
		t.Skipf("Skipping integration test: database not available: %v", err)
	}
	return database
}

func cleanupTestTable(t *testing.T, database *sql.DB) {
	// Очищаем таблицу products для чистоты тестов
	_, err := database.Exec("TRUNCATE TABLE products")
	if err != nil {
		t.Logf("Warning: could not truncate products table: %v", err)
	}
}

func TestIntegration_CreateAndRetrieve(t *testing.T) {
	database := getTestDB(t)
	repo := productRepo.NewPostgresRepository(database)
	service := New(repo)
	defer cleanupTestTable(t, database)

	product := &entity.Product{
		Name:        "Integration Test Product",
		Description: "Testing integration with real DB",
		Price:       99.99,
		Quantity:    100,
	}

	id, err := service.Create(product)
	if err != nil {
		t.Fatalf("Failed to create product: %v", err)
	}

	retrieved, err := service.GetByID(id)
	if err != nil {
		t.Fatalf("Failed to retrieve product: %v", err)
	}

	if retrieved.Name != "Integration Test Product" {
		t.Errorf("Expected name 'Integration Test Product', got %s", retrieved.Name)
	}

	if retrieved.Price != 99.99 {
		t.Errorf("Expected price 99.99, got %f", retrieved.Price)
	}

	if retrieved.Quantity != 100 {
		t.Errorf("Expected quantity 100, got %d", retrieved.Quantity)
	}
}

func TestIntegration_UpdateProduct(t *testing.T) {
	database := getTestDB(t)
	repo := productRepo.NewPostgresRepository(database)
	service := New(repo)
	defer cleanupTestTable(t, database)

	product := &entity.Product{
		Name:     "Original Name",
		Price:    50.0,
		Quantity: 10,
	}

	id, err := service.Create(product)
	if err != nil {
		t.Fatalf("Failed to create product: %v", err)
	}

	updated := &entity.Product{
		Name:     "Updated Name",
		Price:    75.0,
		Quantity: 20,
	}

	err = service.Update(id, updated)
	if err != nil {
		t.Fatalf("Failed to update product: %v", err)
	}

	retrieved, err := service.GetByID(id)
	if err != nil {
		t.Fatalf("Failed to retrieve product: %v", err)
	}

	if retrieved.Name != "Updated Name" {
		t.Errorf("Expected name 'Updated Name', got %s", retrieved.Name)
	}

	if retrieved.Price != 75.0 {
		t.Errorf("Expected price 75.0, got %f", retrieved.Price)
	}
}

func TestIntegration_DeleteProduct(t *testing.T) {
	database := getTestDB(t)
	repo := productRepo.NewPostgresRepository(database)
	service := New(repo)
	defer cleanupTestTable(t, database)

	product := &entity.Product{
		Name:     "To Delete",
		Price:    25.0,
		Quantity: 5,
	}

	id, err := service.Create(product)
	if err != nil {
		t.Fatalf("Failed to create product: %v", err)
	}

	err = service.Delete(id)
	if err != nil {
		t.Fatalf("Failed to delete product: %v", err)
	}

	_, err = service.GetByID(id)
	if err == nil {
		t.Error("Expected error when retrieving deleted product, got nil")
	}
}

func TestIntegration_GetAll(t *testing.T) {
	database := getTestDB(t)
	repo := productRepo.NewPostgresRepository(database)
	service := New(repo)
	defer cleanupTestTable(t, database)

	// Создаем несколько продуктов
	for i := 1; i <= 5; i++ {
		product := &entity.Product{
			Name:     fmt.Sprintf("Product %d", i),
			Price:    float64(i*10) + 0.99,
			Quantity: i * 10,
		}
		_, err := service.Create(product)
		if err != nil {
			t.Fatalf("Failed to create product %d: %v", i, err)
		}
	}

	products, err := service.GetAll()
	if err != nil {
		t.Fatalf("Failed to get all products: %v", err)
	}

	if len(products) != 5 {
		t.Errorf("Expected 5 products, got %d", len(products))
	}
}
