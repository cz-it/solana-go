package transaction

type accountMeta struct {
	pubkey     PublicKey
	isSigner   bool
	isWritable bool
}

// TransactionInstruction is instruction for transaction
type TransactionInstruction struct {
	Keys      []accountMeta
	ProgramID PublicKey
	data      []byte
}

type signatureInfo struct {
	signature []byte
	publicKey PublicKey
}

type nonceInformation struct {
	nonce            string
	nonceInstruction TransactionInstruction
}

// Transaction is a transaction to send to chain
type Transaction struct {
	Signatures      []*signatureInfo
	Instructions    []*TransactionInstruction
	RecentBlockhash string
	NonceInfo       nonceInformation
}

func (t *Transaction) Add(instruction *TransactionInstruction) error {
	return nil
}
