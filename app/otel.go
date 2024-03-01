package app

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// Creates a trace provider that will be used globally by the application
func InitTracer() (*sdktrace.TracerProvider, error) {

	exporter, err := otlptracegrpc.New(
		context.Background(),
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint("127.0.0.1:4317"),
	)
	if err != nil {
		return nil, err
	}

	//Create a new trace provider with the exporter
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter, sdktrace.WithBatchTimeout(5)),
	)
	//Configure {tp}  as the global trace provider
	otel.SetTracerProvider(tp)
	//Configure the global propagator to use the W3C Trace Context
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}),
	)

	//Return the trace provider
	return tp, nil
}
