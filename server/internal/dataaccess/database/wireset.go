package database

import (
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	ProvideDatabase,
	// InitializeDB,
	// InitializeGoquDB,
	NewAccountAccessor,
	NewAccountPasswordDataAccessor,
)
