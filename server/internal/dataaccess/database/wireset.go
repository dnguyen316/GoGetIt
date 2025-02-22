package database

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	ProvideDatabase,
	wire.Bind(new(Database), new(*goqu.Database)),
	// InitializeDB,
	// InitializeGoquDB,
	NewAccountAccessor,
	NewAccountPasswordDataAccessor,
)
