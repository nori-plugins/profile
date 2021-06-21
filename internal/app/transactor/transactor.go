package transactor

import (
	"github.com/google/wire"
	"github.com/nori-plugins/profile/internal/transactor"
)

var TransactorSet = wire.NewSet(
	wire.Struct(new(transactor.Params), "Db", "Logger"),
	transactor.New)
