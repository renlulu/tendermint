package p2p

import (
	cryptoAmino "github.com/renlulu/tendermint/crypto/encoding/amino"
	"github.com/tendermint/go-amino"
)

var cdc = amino.NewCodec()

func init() {
	cryptoAmino.RegisterAmino(cdc)
}
