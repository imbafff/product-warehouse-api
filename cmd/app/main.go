package main

import (
	"log"

	httpDelivery "github.com/imbafff/product-warehouse-api/internal/delivery/http"
	"github.com/imbafff/product-warehouse-api/internal/delivery/http/handler"
	"github.com/imbafff/product-warehouse-api/internal/infrastructure/config"
	"github.com/imbafff/product-warehouse-api/internal/infrastructure/db"
	productRepo "github.com/imbafff/product-warehouse-api/internal/repository/product"
	productUC "github.com/imbafff/product-warehouse-api/internal/usecase/product"
)

func main() {
	cfg := config.Load()

	database, err := db.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal("failed to connect to db:", err)
	}

	repo := productRepo.NewPostgresRepository(database)
	usecase := productUC.New(repo)
	h := handler.NewProductHandler(usecase)

	r := httpDelivery.NewRouter(h)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("failed to run server:", err)
	}
}
