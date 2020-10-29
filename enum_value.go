package solana

// CommitmentValue commitment
type CommitmentValue string

// CommitmentValue commitment
const (
	CommitmentRecent       = "recent"
	CommitmentMax          = "max"
	CommitmentRoot         = "root"
	CommitmentSingleGossip = "singleGossip"
)

// ResponseEncoding is encoding
type ResponseEncoding string

// ResponseEncoding is encoding
const (
	ResponseEncodingBase64     = "base64"
	ResponseEncodingBase58     = "base58"
	ResponseEncodingJSON       = "json"
	ResponseEncodingJSONParsed = "jsonParsed"
)

// AccountType type for account
type AccountType string

// AccountType type for account
const (
	AccountTypeCirculating    = "circulating"
	AccountTypeNonCirculating = "nonCirculating"
)
