package router

import (
	"github.com/Alex-omosa/pricing-service/handler"
	"github.com/Alex-omosa/pricing-service/services"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

type Router struct {
	Tracer  trace.Tracer
	handler handler.Handler
}

func NewRouter(
	tracer trace.Tracer,
) *Router {
	return &Router{
		Tracer:  tracer,
		handler: *handler.NewHandler(tracer, services.NewPricingService(tracer)),
	}
}

func (r *Router) AddRoutes(router *gin.Engine) {

	trips := router.Group("/api/v1/")
	{
		trips.GET("/prices", r.handler.GetAllPrices)
	}
	driver := router.Group("/api/v1/")
	{
		driver.GET("/drivers", r.handler.GetPrice)
	}
}
