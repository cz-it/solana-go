package solana

import (
	"github.com/kevinburke/nacl/sign"
)

// Account is a account object with Public Key
type Account struct {
	publicKey  [32]byte
	privateKey [64]byte
}

// NewAccount create a account with random seed
func NewAccount() *Account {
	a := &Account{}
	pubKey, priKey, err := sign.Keypair(nil)
	if err != nil {
		return nil
	}
	copy(a.publicKey[:], pubKey)
	copy(a.privateKey[:], priKey)
	return a
}

// NewAccountWithSecretKey create an account with secretKey
func NewAccountWithSecretKey(secretKey [64]byte) *Account {
	a := &Account{}
	copy(a.privateKey[:], secretKey[:])
	prikey := sign.PrivateKey(secretKey[:])
	prikey.Public()
	pubkey := prikey.Public().(sign.PublicKey)
	copy(a.publicKey[:], pubkey)
	return a
}

// PublicKey return a PublicKey object for account
func (a *Account) PublicKey() *PublicKey {
	p, err := NewPublicKeyWithKey(a.publicKey)
	if err != nil {
		return nil
	}
	return p
}

// SecretKey return  **unencrypted** secret key for this account
func (a *Account) SecretKey() [64]byte {
	return a.privateKey
}
