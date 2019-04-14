package monitor

import (
	ctypes "github.com/renlulu/tendermint/rpc/core/types"
	"github.com/tendermint/go-amino"
)

var cdc = amino.NewCodec()

func init() {
	ctypes.RegisterAmino(cdc)
}
