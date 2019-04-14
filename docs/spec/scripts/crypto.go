package main

import (
	"fmt"
	"os"

	cryptoAmino "github.com/renlulu/tendermint/crypto/encoding/amino"
	"github.com/tendermint/go-amino"
)

func main() {
	cdc := amino.NewCodec()
	cryptoAmino.RegisterAmino(cdc)
	cdc.PrintTypes(os.Stdout)
	fmt.Println("")
}
