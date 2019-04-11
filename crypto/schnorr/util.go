package schnorr

import (
	"bytes"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

func Sha256(data []byte) []byte {
	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}

func Compress(curve elliptic.Curve, x, y *big.Int) []byte {
	return Marshal(curve, x, y, true)
}

func Marshal(curve elliptic.Curve, x, y *big.Int, compress bool) []byte {
	byteLen := (curve.Params().BitSize + 7) >> 3

	if compress {
		ret := make([]byte, 1+byteLen)
		if y.Bit(0) == 0 {
			ret[0] = 2
		} else {
			ret[0] = 3
		}
		xBytes := x.Bytes()
		copy(ret[1+byteLen-len(xBytes):], xBytes)
		return ret
	}

	ret := make([]byte, 1+2*byteLen)
	ret[0] = 4 // uncompressed point
	xBytes := x.Bytes()
	copy(ret[1+byteLen-len(xBytes):], xBytes)
	yBytes := y.Bytes()
	copy(ret[1+2*byteLen-len(yBytes):], yBytes)
	return ret
}

func bigIntToBytes(bi *big.Int) []byte {
	b1, b2 := [32]byte{}, bi.Bytes()
	copy(b1[32-len(b2):], b2)
	return b1[:]
}

func hash(Q []byte, pubKey []byte, msg []byte) []byte {
	var buffer bytes.Buffer
	buffer.Write(Q)
	buffer.Write(pubKey[:33])
	buffer.Write(msg)
	return Sha256(buffer.Bytes())
}


func GenerateDRN(nonce []byte) ([]byte, error) {
	var buffer bytes.Buffer
	buffer.Write(generateRandomBytes(32))
	buffer.WriteString("Schnorr+SHA256  ")
	hmacDRBG := NewHmacDRBG(generateRandomBytes(32), nonce, buffer.Bytes())

	return hmacDRBG.Generate(int32(32), []byte{})
}

func generateRandomBytes(size int32) []byte {
	randomBytes := make([]byte, size)
	_, _ = rand.Read(randomBytes)
	return randomBytes
}
