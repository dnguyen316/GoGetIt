//go:build wireinject
// +build wireinject

//
//go:generate go run github.com/google/wire/cmd/wire

package wiring

import (
	"github.com/dnguyen316/GoGetIt/server/internal/config"
	"github.com/dnguyen316/GoGetIt/server/internal/dataaccess"
	"github.com/dnguyen316/GoGetIt/server/internal/handler"
	"github.com/dnguyen316/GoGetIt/server/internal/handler/grpc"
	"github.com/dnguyen316/GoGetIt/server/internal/logic"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	config.WireSet,
	logic.WireSet,
	handler.WireSet,
	dataaccess.WireSet,
)

func InitializeGRPCServer(configFilePath config.ConfigFilePath) (grpc.Server, func(), error) {
	wire.Build(WireSet)

	return nil, nil, nil
}
