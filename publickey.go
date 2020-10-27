package solana

import (
	"errors"
	"math/big"

	"github.com/btcsuite/btcutil/base58"
)

// PublicKey is a PublicKey object
type PublicKey struct {
	bn big.Int
}

// ToBase58 return the base-58 representation of the public key
func (pk *PublicKey) ToBase58() string {
	return base58.Encode(pk.bn.Bytes())
}

// ToString return a string representation of the public key
func (pk *PublicKey) ToString() string {
	return pk.ToBase58()
}

// NewPublicKey create a PublicKey object with a key of base58 string
func NewPublicKey(key string) (publicKey *PublicKey, err error) {
	decode := base58.Decode(key)
	if len(decode) != 32 {
		return nil, errors.New("Invalid public key input")
	}
	publicKey = &PublicKey{}
	publicKey.bn.SetBytes(decode)
	return publicKey, nil
}

// NewPublicKeyWithKey create a Publickey object with a [32]byte pulbic key
func NewPublicKeyWithKey(key [32]byte) (publicKey *PublicKey, err error) {
	publicKey = &PublicKey{}
	publicKey.bn.SetBytes(key[:])
	return publicKey, nil
}
