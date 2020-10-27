package solana

import (
	"testing"
)

type accountTestCase struct {
	secretKey [64]byte
	base58    string
}

var accountCases = []accountTestCase{
	{
		[64]byte{255, 20, 128, 9, 125, 112, 25, 86, 79, 228, 38, 92, 13, 193, 154, 239, 240, 56, 8, 237, 39, 181, 124, 42, 219, 193, 220, 115, 113, 104, 190, 117, 244, 78, 138, 70, 192, 140, 192, 159, 141, 60, 97, 101, 12, 213, 132, 163, 141, 41, 6, 112, 158, 115, 112, 49, 153, 102, 4, 214, 253, 14, 33, 131},
		"HSfwVfB7RUF1SKCd4yrz8KZp7TU262Y5BeZZN1tdCTVk",
	},
}

func TestNewAccountWithKey(t *testing.T) {
	for _, c := range accountCases {
		skey := c.secretKey
		account := NewAccountWithSecretKey(skey)
		if account.PublicKey().ToBase58() != c.base58 {
			t.Errorf("base58 publick key (%s)is not equal", account.PublicKey().ToBase58())
		}
	}
}
