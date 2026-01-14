package http

import (
	"github.com/imbafff/product-warehouse-api/internal/delivery/http/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(h *handler.ProductHandler) *gin.Engine {
	r := gin.Default()

	products := r.Group("/products")
	{
		products.POST("", h.Create)
		products.GET("", h.GetAll)
		products.GET("/:id", h.GetByID)
		products.PUT("/:id", h.Update)
		products.DELETE("/:id", h.Delete)
	}

	return r
}
