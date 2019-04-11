package secp256k1

import (
	"github.com/renlulu/tendermint/crypto/schnorr"
)

// Sign creates an ECDSA signature on curve Secp256k1, using SHA256 on the msg.
//func (privKey PrivKeySecp256k1) Sign(msg []byte) ([]byte, error) {
//	rsv, err := secp256k1.Sign(crypto.Sha256(msg), privKey[:])
//	if err != nil {
//		return nil, err
//	}
//	// we do not need v  in r||s||v:
//	rs := rsv[:len(rsv)-1]
//	return rs, nil
//}
//
//func (pubKey PubKeySecp256k1) VerifyBytes(msg []byte, sig []byte) bool {
//	return secp256k1.VerifySignature(pubKey[:], crypto.Sha256(msg), sig)
//}

func (privKey PrivKeySecp256k1) Sign(msg []byte) ([]byte, error) {
	rs := [64]byte{}
	pubKeyI := privKey.PubKey()
	pubKeyArray := pubKeyI.(PubKeySecp256k1)

	puk := [33]byte(pubKeyArray)

	k, _ := schnorr.GenerateDRN(msg)
	r, s, err := schnorr.TrySign(privKey[:], puk[:], msg, k)
	if err != nil {
		return rs[:], err
	}
	copy(rs[0:32], r)
	copy(rs[32:64], s)
	return rs[:], nil
}

func (pubKey PubKeySecp256k1) VerifyBytes(msg []byte, sig []byte) bool {
	r := sig[:32]
	s := sig[32:]
	return schnorr.Verify(pubKey[:], msg, r, s)
}
