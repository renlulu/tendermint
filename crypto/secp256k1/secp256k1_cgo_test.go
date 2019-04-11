package secp256k1

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestPrivKeySecp256k1_Sign(t *testing.T) {

	privBytes := [32]byte{}
	copy(privBytes[:], hex_bytes("0F494B8312E8D257E51730C78F8FE3B47B6840C59AAAEC7C2EBE404A2DE8B25A"))
	privateKey := PrivKeySecp256k1(privBytes)

	msg := hex_bytes("A7F1D92A82C8D8FE434D98558CE2B347171198542F112D0558F56BD68807999248336241F30D23E55F30D1C8ED610C4B0235398184B814A29CB45A672ACAE548E9C5F1B0C4158AE59B4D39F6F7E8A105D3FEEDA5D5F3D9E45BFA6CC351E220AE0CE106986D61FF34A11E19FD3650E9B7818FC33A1E0FC02C44557AC8AB50C9B2DEB2F6B5E24C4FDD9F8867BDCE1FF261008E7897970E346207D75E47A158298E5BA2F56246869CC42E362A02731264E60687EF5309D108534F51F8658FB4F080B7CB19EE9AEBD718CC4FA27C8C37DFC1ADA5D133D13ABE03F021E9B1B78CCBD82F7FF2B38C6D48D01E481B2D4FAF7171805FD7F2D39EF4C4F19B9496E81DAB8193B3737E1B27D9C43957166441B93515E8F03C95D8E8CE1E1864FAAD68DDFC5932130109390B0F1FE5CA716805F8362E98DCCAADC86ADBED25801A9A9DCFA6264319DDAFE83A89C51F3C6D199D38DE10E660C37BE872C3F2B31660DE8BC95902B9103262CDB941F77376F5D3DBB7A3D5A387797FC4819A035ECA704CEDB37110EE7F206B0C8805AAEBF4963E7C4708CE8D4E092366E71792A8A3B2BBCDEE321B3E15380C541EF0930888969F7457AFE18588826A419D58311C1784B5484EECDB393F6A0ACA11B91DF0866B500B8DEE501FD7EB9BCE09A17D74124B4605ADFC0777BED9816D8D7E8488544A18D8045CB3283B0A752B881B5F500FADB59010E63D")
	rs, err := privateKey.Sign(msg)

	if err != nil {
		t.Error(err.Error())
	}

	sigString := hex.EncodeToString(rs)
	fmt.Printf("sig = %s\n", sigString)

	verify := privateKey.PubKey().VerifyBytes(msg, rs)

	fmt.Printf("verify = %t", verify)

	fakeMsg := hex_bytes("B7F1D92A82C8D8FE434D98558CE2B347171198542F112D0558F56BD68807999248336241F30D23E55F30D1C8ED610C4B0235398184B814A29CB45A672ACAE548E9C5F1B0C4158AE59B4D39F6F7E8A105D3FEEDA5D5F3D9E45BFA6CC351E220AE0CE106986D61FF34A11E19FD3650E9B7818FC33A1E0FC02C44557AC8AB50C9B2DEB2F6B5E24C4FDD9F8867BDCE1FF261008E7897970E346207D75E47A158298E5BA2F56246869CC42E362A02731264E60687EF5309D108534F51F8658FB4F080B7CB19EE9AEBD718CC4FA27C8C37DFC1ADA5D133D13ABE03F021E9B1B78CCBD82F7FF2B38C6D48D01E481B2D4FAF7171805FD7F2D39EF4C4F19B9496E81DAB8193B3737E1B27D9C43957166441B93515E8F03C95D8E8CE1E1864FAAD68DDFC5932130109390B0F1FE5CA716805F8362E98DCCAADC86ADBED25801A9A9DCFA6264319DDAFE83A89C51F3C6D199D38DE10E660C37BE872C3F2B31660DE8BC95902B9103262CDB941F77376F5D3DBB7A3D5A387797FC4819A035ECA704CEDB37110EE7F206B0C8805AAEBF4963E7C4708CE8D4E092366E71792A8A3B2BBCDEE321B3E15380C541EF0930888969F7457AFE18588826A419D58311C1784B5484EECDB393F6A0ACA11B91DF0866B500B8DEE501FD7EB9BCE09A17D74124B4605ADFC0777BED9816D8D7E8488544A18D8045CB3283B0A752B881B5F500FADB59010E63D")

	fakeVerify := privateKey.PubKey().VerifyBytes(fakeMsg, rs)

	fmt.Printf("fake verify = %t\n", fakeVerify)

}

func hex_bytes(hs string) []byte {
	data, err := hex.DecodeString(hs)
	if err != nil {
		panic("cannot convert hex string to byte array\n")
	}
	return data
}
