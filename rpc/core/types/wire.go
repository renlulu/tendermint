package core_types

import (
	"github.com/renlulu/tendermint/types"
	"github.com/tendermint/go-amino"
)

func RegisterAmino(cdc *amino.Codec) {
	types.RegisterEventDatas(cdc)
	types.RegisterBlockAmino(cdc)
}
