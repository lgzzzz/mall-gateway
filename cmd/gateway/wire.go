//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"

	"gateway/internal/client"
	"mall-gateway/internal/conf"
	"mall-gateway/internal/server"
	"mall-gateway/internal/service"
)

func wireApp(*conf.Config, log.Logger, trace.Tracer) (*kratos.App, func(), error) {
	panic(wire.Build(
		client.ProviderSet,
		service.ProviderSet,
		server.ProviderSet,
		newApp,
	))
}
