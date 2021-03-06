// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"web3/app/shop/service/internal/biz"
	"web3/app/shop/service/internal/conf"
	"web3/app/shop/service/internal/data"
	"web3/app/shop/service/internal/server"
	"web3/app/shop/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, discovery *conf.Discovery, confData *conf.Data, logger log.Logger, tracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	dataData, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	shopRepo := data.NewShopRepo(dataData, logger)
	registryDiscovery := server.NewDiscovery(discovery)
	userClient := server.NewUserServiceClient(registryDiscovery)
	shopUseCase := biz.NewShopUseCase(shopRepo, logger, userClient)
	shopService := service.NewShopService(shopUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, logger, tracerProvider, shopService)
	grpcServer := server.NewGRPCServer(confServer, logger, tracerProvider, shopService)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
	}, nil
}
