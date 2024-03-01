package services

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

type PricingService struct {
	tracer trace.Tracer
}

func NewPricingService(tracer trace.Tracer) *PricingService {
	return &PricingService{
		tracer: tracer,
	}
}

func (p *PricingService) GetAllPrices(ctx context.Context) error {
	_, span := p.tracer.Start(ctx, "GetAllPrices-Service")
	defer span.End()

	return nil
}

func (p *PricingService) GetPrice(ctx context.Context) error {
	_, span := p.tracer.Start(ctx, "GetPrice-Service")
	defer span.End()
	return nil
}
