package handler

import (
	"github.com/Alex-omosa/pricing-service/services"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

type Handler struct {
	tracer  trace.Tracer
	service *services.PricingService
}

func NewHandler(tracer trace.Tracer,
	service *services.PricingService) *Handler {
	return &Handler{
		tracer:  tracer,
		service: service,
	}
}

func (h *Handler) GetAllPrices(c *gin.Context) {
	ctx, span := h.tracer.Start(c.Request.Context(), "GetAllPrices")
	defer span.End()

	err := h.service.GetAllPrices(ctx)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Get all trips",
	})
}

func (h *Handler) GetPrice(c *gin.Context) {
	ctx, span := h.tracer.Start(c.Request.Context(), "GetPrice")
	defer span.End()

	err := h.service.GetPrice(ctx)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
}
