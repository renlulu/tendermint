package blockchain

import (
	amino "github.com/tendermint/go-amino"
	"github.com/renlulu/tendermint/types"
)

var cdc = amino.NewCodec()

func init() {
	RegisterBlockchainMessages(cdc)
	types.RegisterBlockAmino(cdc)
}
