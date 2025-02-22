package handler

import (
	"github.com/dnguyen316/GoGetIt/server/internal/handler/grpc"
	"github.com/dnguyen316/GoGetIt/server/internal/handler/http"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	grpc.WireSet,
	http.WireSet,
)
