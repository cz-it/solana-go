package solana

import (
	"testing"
)

type publicKeyTestCase struct {
	publicKey string
	bn        string
}

var publicKeyCases = []publicKeyTestCase{
	{"HSfwVfB7RUF1SKCd4yrz8KZp7TU262Y5BeZZN1tdCTVk", "110503103473839698732642032862719776302565871678930368569016380966738594767235"},
}

func TestNewPublicKey(t *testing.T) {
	for _, c := range publicKeyCases {
		p, err := NewPublicKey(c.publicKey)
		if err != nil {
			t.Errorf("new Publick key error:%s", err.Error())
		}
		if p.bn.String() != c.bn {
			t.Errorf("bn.String(%s) is not equal", p.bn.String())
		}
	}
}

func TestBase58(t *testing.T) {
	for _, c := range publicKeyCases {
		p, err := NewPublicKey(c.publicKey)
		if err != nil {
			t.Errorf("new Publick key error:%s", err.Error())
		}
		if p.ToBase58() != c.publicKey {
			t.Errorf("ToBase58(%s) is not equal", p.ToBase58())
		}
	}
}
