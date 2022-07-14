package main

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/ujunglangit-id/go-gin-boilerplate/internal/model/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"time"
)

func tracerProvider(cfg *config.Config) (*tracesdk.TracerProvider, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(cfg.Monitoring.TracingUrl)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(cfg.Monitoring.TracingAppName),
			attribute.String("environment", cfg.Server.Environment),
		)),
	)
	return tp, nil
}

func main() {
	// init logging
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// init config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal().Err(err)
	}
	// init open telemetry
	tp, err := tracerProvider(cfg)
	if err != nil {
		log.Fatal().Err(err)
	}

	otel.SetTracerProvider(tp)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Cleanly shutdown and flush telemetry when the application exits.
	defer func(ctx context.Context) {
		// Do not make the application hang when it is shutdown.
		ctx, cancel = context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatal().Err(err)
		}
	}(ctx)

	// init repository

	// init usecase

	// init delivery

	// initiate graceful shutdown

	// start http server
}
