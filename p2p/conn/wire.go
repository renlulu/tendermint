package conn

import (
	cryptoAmino "github.com/renlulu/tendermint/crypto/encoding/amino"
	"github.com/tendermint/go-amino"
)

var cdc *amino.Codec = amino.NewCodec()

func init() {
	cryptoAmino.RegisterAmino(cdc)
	RegisterPacket(cdc)
}
