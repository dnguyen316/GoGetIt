package dataaccess

import (
	"github.com/dnguyen316/GoGetIt/server/internal/dataaccess/database"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	database.WireSet,
)
