package types

import (
	"github.com/renlulu/tendermint/types"
	"github.com/tendermint/go-amino"
)

var cdc = amino.NewCodec()

func init() {
	types.RegisterBlockAmino(cdc)
}
