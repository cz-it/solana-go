package solana

import (
	"fmt"
)

// RPCResult  is result for JSON RPC
type RPCResult struct {
	Context struct {
		Slot int `json:"slot"`
	} `json:"context"`
	Value interface{} `josn:"value,omitempty"`
}

// CommitmentConfig is commitment
type CommitmentConfig struct {
	Commitment string `json:"commitment"`
}

// EpochInfo is response of getEpochInfo
type EpochInfo struct {
	AbsoluteSlot int `json:"absoluteSlot"`
	BlockHeight  int `json:"blockHeight"`
	Epoch        int `json:"epoch"`
	SlotIndex    int `json:"slotIndex"`
	SlotsInEpoch int `json:"slotsInEpoch"`
}

// GetEpochInfo Fetch the Epoch Info parameters
func (c *Connection) GetEpochInfo(commitment CommitmentValue) (*EpochInfo, error) {
	var result EpochInfo
	cmm := CommitmentConfig{
		Commitment: string(commitment),
	}
	err := c.client.Call(&result, "getEpochInfo", cmm)
	if err != nil {
		return nil, fmt.Errorf("GetEpochInfo with error %s", err.Error())
	}
	return &result, nil
}

// EpochSchedule is response of getEpochSchedule
type EpochSchedule struct {
	FirstNormalEpoch         int  `json:"firstNormalEpoch"`
	FirstNormalSlot          int  `json:"firstNormalSlot"`
	LeaderScheduleSlotOffset int  `json:"leaderScheduleSlotOffset"`
	SlotsPerEpoch            int  `json:"slotsPerEpoch"`
	Warmup                   bool `json:"warmup"`
}

// GetEpochSchedule Fetch the Epoch Schedule parameters
func (c *Connection) GetEpochSchedule() (*EpochSchedule, error) {
	var result EpochSchedule
	err := c.client.Call(&result, "getEpochSchedule")
	if err != nil {
		return nil, fmt.Errorf("GetEpochSchedule with error %s", err.Error())
	}
	return &result, nil
}

// GetAccountInfoOpts is optional for GetAccountInfo
type GetAccountInfoOpts struct {
	Encoding   string `json:"encoding"`
	Commitment string `json:"commitment"`
	DataSlice  *struct {
		Offset int `json:"offset"`
		Length int `json:"length"`
	} `json:"dataSlice,omitempty"`
}

// AccountInfo is response of getAccountInfo
type AccountInfo struct {
	Data       []string `json:"data"`
	Executable bool     `json:"executable"`
	Lamports   int64    `json:"lamports"`
	Owner      string   `json:"owner"`
	RentEpoch  int      `json:"rentEpoch"`
}

// GetAccountInfo Fetch all the account info for the specified public key, return with context
func (c *Connection) GetAccountInfo(publicKey string, opts *GetAccountInfoOpts) (*AccountInfo, error) {
	var result RPCResult
	var info AccountInfo
	result.Value = &info
	err := c.client.Call(&result, "getAccountInfo", publicKey, opts)
	if err != nil {
		return nil, fmt.Errorf("GetAccountInfo with error %s", err.Error())
	}
	return &info, nil
}

// GetBalance Fetch the balance for the specified public key, return with context
func (c *Connection) GetBalance(publicKey string, commitment CommitmentValue) (int, error) {
	//var result BalanceInfo
	var result RPCResult
	cmm := CommitmentConfig{
		Commitment: string(commitment),
	}
	var value int
	result.Value = &value
	err := c.client.Call(&result, "getBalance", publicKey, cmm)
	if err != nil {
		return 0, fmt.Errorf("GetBalance with error %s", err.Error())
	}
	return value, nil
}

// BlockCommitmentInfo is response for getBlockCommitment
type BlockCommitmentInfo struct {
	Commitment []int `json:"commitment"`
	TotalStake int   `json:"totalStake"`
}

// GetBlockCommitment Fetch block commitment for solt (block)
func (c *Connection) GetBlockCommitment(solt uint64) (*BlockCommitmentInfo, error) {
	var result BlockCommitmentInfo
	err := c.client.Call(&result, "getBlockCommitment", solt)
	if err != nil {
		return nil, fmt.Errorf("GetBlockCommitment with error %s", err.Error())
	}
	return &result, nil
}

// GetBlockTime Fetch the estimated production time of a block
func (c *Connection) GetBlockTime(solt uint64) (uint64, error) {
	var result uint64
	err := c.client.Call(&result, "getBlockTime", solt)
	if err != nil {
		return 0, fmt.Errorf("GetBlockTime with error %s", err.Error())
	}
	return result, nil
}

// ClusterNodeInfo is response for getClusterNodes
type ClusterNodeInfo struct {
	Gossip  string `json:"gossip"`
	Pubkey  string `json:"pubkey"`
	RPC     string `json:"rpc"`
	Tpu     string `json:"tpu"`
	Version string `json:"version"`
}

// GetClusterNodes Return the list of nodes that are currently participating in the cluster
func (c *Connection) GetClusterNodes() ([]ClusterNodeInfo, error) {
	var result []ClusterNodeInfo
	err := c.client.Call(&result, "getClusterNodes")
	if err != nil {
		return nil, fmt.Errorf("GetClusterNodes with error %s", err.Error())
	}
	return result, nil
}

// ConfirmedBlockInfo is response for getConfirmedBlock
type ConfirmedBlockInfo struct {
	BlockTime         interface{}   `json:"blockTime"`
	Blockhash         string        `json:"blockhash"`
	ParentSlot        int           `json:"parentSlot"`
	PreviousBlockhash string        `json:"previousBlockhash"`
	Rewards           []interface{} `json:"rewards"`
	Transactions      []struct {
		Meta        interface{} `json:"meta"`
		Transaction struct {
			Message struct {
				AccountKeys []struct {
					Pubkey   string `json:"pubkey"`
					Signer   bool   `json:"signer"`
					Writable bool   `json:"writable"`
				} `json:"accountKeys"`
				Instructions []struct {
					Accounts  []string `json:"accounts"`
					Data      string   `json:"data"`
					ProgramID string   `json:"programId"`
				} `json:"instructions"`
				RecentBlockhash string `json:"recentBlockhash"`
			} `json:"message"`
			Signatures []string `json:"signatures"`
		} `json:"transaction"`
	} `json:"transactions"`
}

// ConfirmedBlockEncodedInfo is response for getConfirmedBlock
type ConfirmedBlockEncodedInfo struct {
	BlockTime         interface{}   `json:"blockTime"`
	Blockhash         string        `json:"blockhash"`
	ParentSlot        int           `json:"parentSlot"`
	PreviousBlockhash string        `json:"previousBlockhash"`
	Rewards           []interface{} `json:"rewards"`
	Transactions      []struct {
		Meta        interface{} `json:"meta"`
		Transaction []string    `json:"transaction"`
	} `json:"transactions"`
}

// ConfirmedBlockJSONInfo is response for getConfirmedBlock
type ConfirmedBlockJSONInfo struct {
	BlockTime         interface{}   `json:"blockTime"`
	Blockhash         string        `json:"blockhash"`
	ParentSlot        int           `json:"parentSlot"`
	PreviousBlockhash string        `json:"previousBlockhash"`
	Rewards           []interface{} `json:"rewards"`
	Transactions      []struct {
		Meta        interface{} `json:"meta"`
		Transaction struct {
			Message struct {
				AccountKeys []string `json:"accountKeys"`
				Header      struct {
					NumReadonlySignedAccounts   int `json:"numReadonlySignedAccounts"`
					NumReadonlyUnsignedAccounts int `json:"numReadonlyUnsignedAccounts"`
					NumRequiredSignatures       int `json:"numRequiredSignatures"`
				} `json:"header"`
				Instructions []struct {
					Accounts       []int  `json:"accounts"`
					Data           string `json:"data"`
					ProgramIDIndex int    `json:"programIdIndex"`
				} `json:"instructions"`
				RecentBlockhash string `json:"recentBlockhash"`
			} `json:"message"`
			Signatures []string `json:"signatures"`
		} `json:"transaction"`
	} `json:"transactions"`
}

// GetConfirmedBlock Fetch a list of Transactions and transaction statuses from the cluster for a confirmed block
func (c *Connection) GetConfirmedBlock(solt uint64, encoding ResponseEncoding) (interface{}, error) {
	var result ConfirmedBlockInfo
	var eresult ConfirmedBlockEncodedInfo
	var jresult ConfirmedBlockJSONInfo
	if encoding == ResponseEncodingJSONParsed {
		err := c.client.Call(&result, "getConfirmedBlock", solt, encoding)
		if err != nil {
			return nil, fmt.Errorf("GetConfirmedBlock with error %s", err.Error())
		}
		return &result, nil
	}

	if encoding == ResponseEncodingJSON {
		err := c.client.Call(&jresult, "getConfirmedBlock", solt, encoding)
		if err != nil {
			return nil, fmt.Errorf("GetConfirmedBlock with error %s", err.Error())
		}
		return &jresult, nil
	}

	if encoding == ResponseEncodingBase58 || encoding == ResponseEncodingBase64 {
		err := c.client.Call(&eresult, "getConfirmedBlock", solt, encoding)
		if err != nil {
			return nil, fmt.Errorf("GetConfirmedBlock with error %s", err.Error())
		}
		return &eresult, nil
	}

	return nil, fmt.Errorf("GetConfirmedBlock with unsupport encoding")
}

// GetConfirmedBlocks Returns a list of confirmed blocks between two slots
func (c *Connection) GetConfirmedBlocks(start, end uint64) ([]uint64, error) {
	var result []uint64
	err := c.client.Call(&result, "getConfirmedBlocks", start, end)
	if err != nil {
		return nil, fmt.Errorf("GetConfirmedBlocks with error %s", err.Error())
	}
	return result, nil
}

// GetConfirmedBlocksWithLimit Returns a list of confirmed blocks starting at the given slot
func (c *Connection) GetConfirmedBlocksWithLimit(start, limit uint64) ([]uint64, error) {
	var result []uint64
	err := c.client.Call(&result, "getConfirmedBlocksWithLimit", start, limit)
	if err != nil {
		return nil, fmt.Errorf("GetConfirmedBlocksWithLimit with error %s", err.Error())
	}
	return result, nil
}

// ConfirmedSignatureInfo is response for getConfirmedSignaturesForAddress2
type ConfirmedSignatureInfo struct {
	Err       interface{} `json:"err"`
	Memo      interface{} `json:"memo"`
	Signature string      `json:"signature"`
	Slot      int         `json:"slot"`
}

// GetConfirmedSignaturesForAddress2Opts is optinal for getConfirmedSignaturesForAddress2
type GetConfirmedSignaturesForAddress2Opts struct {
	Limit  int    `json:"limit,omitempty"`
	Before string `json:"before,omitempty"`
	Until  string `json:"until,omitempty"`
}

// GetConfirmedSignaturesForAddress2 Returns confirmed signatures for
//transactions involving an address backwards in time from the provided
//signature or most recent confirmed block
func (c *Connection) GetConfirmedSignaturesForAddress2(publicKey string,
	opts GetConfirmedSignaturesForAddress2Opts) ([]ConfirmedSignatureInfo, error) {
	var result []ConfirmedSignatureInfo
	err := c.client.Call(&result, "getConfirmedSignaturesForAddress2", publicKey, opts)
	if err != nil {
		return nil, fmt.Errorf("GetConfirmedSignaturesForAddress2 with error %s", err.Error())
	}
	return result, nil

}

type ConfirmedTransactionJSONInfo struct {
	Meta struct {
		Err               interface{}   `json:"err"`
		Fee               int           `json:"fee"`
		InnerInstructions []interface{} `json:"innerInstructions"`
		LogMessages       []string      `json:"logMessages"`
		PostBalances      []interface{} `json:"postBalances"`
		PreBalances       []interface{} `json:"preBalances"`
		Status            struct {
			Ok interface{} `json:"Ok"`
		} `json:"status"`
	} `json:"meta"`
	Slot        int `json:"slot"`
	Transaction struct {
		Message struct {
			AccountKeys []string `json:"accountKeys"`
			Header      struct {
				NumReadonlySignedAccounts   int `json:"numReadonlySignedAccounts"`
				NumReadonlyUnsignedAccounts int `json:"numReadonlyUnsignedAccounts"`
				NumRequiredSignatures       int `json:"numRequiredSignatures"`
			} `json:"header"`
			Instructions []struct {
				Accounts       []int  `json:"accounts"`
				Data           string `json:"data"`
				ProgramIDIndex int    `json:"programIdIndex"`
			} `json:"instructions"`
			RecentBlockhash string `json:"recentBlockhash"`
		} `json:"message"`
		Signatures []string `json:"signatures"`
	} `json:"transaction"`
}

type ConfirmedTransactionInfo struct {
	Meta struct {
		Err               interface{}   `json:"err"`
		Fee               int           `json:"fee"`
		InnerInstructions []interface{} `json:"innerInstructions"`
		LogMessages       []string      `json:"logMessages"`
		PostBalances      []interface{} `json:"postBalances"`
		PreBalances       []interface{} `json:"preBalances"`
		Status            struct {
			Ok interface{} `json:"Ok"`
		} `json:"status"`
	} `json:"meta"`
	Slot        int `json:"slot"`
	Transaction struct {
		Message struct {
			AccountKeys []struct {
				Pubkey   string `json:"pubkey"`
				Signer   bool   `json:"signer"`
				Writable bool   `json:"writable"`
			} `json:"accountKeys"`
			Instructions []struct {
				Parsed struct {
					Info struct {
						Amount      string `json:"amount"`
						Authority   string `json:"authority"`
						Destination string `json:"destination"`
						Source      string `json:"source"`
					} `json:"info"`
					Type string `json:"type"`
				} `json:"parsed"`
				Program   string `json:"program"`
				ProgramID string `json:"programId"`
			} `json:"instructions"`
			RecentBlockhash string `json:"recentBlockhash"`
		} `json:"message"`
		Signatures []string `json:"signatures"`
	} `json:"transaction"`
}

type ConfirmedTransactionEncodedInfo struct {
	Meta struct {
		Err               interface{}   `json:"err"`
		Fee               int           `json:"fee"`
		InnerInstructions []interface{} `json:"innerInstructions"`
		LogMessages       []string      `json:"logMessages"`
		PostBalances      []interface{} `json:"postBalances"`
		PreBalances       []interface{} `json:"preBalances"`
		Status            struct {
			Ok interface{} `json:"Ok"`
		} `json:"status"`
	} `json:"meta"`
	Slot        int      `json:"slot"`
	Transaction []string `json:"transaction"`
}

// GetConfirmedTransaction Returns transaction details for a confirmed transaction
func (c *Connection) GetConfirmedTransaction(publicKey string, encoding ResponseEncoding) (interface{}, error) {
	var result ConfirmedTransactionInfo
	var eresult ConfirmedTransactionEncodedInfo
	var jresult ConfirmedTransactionJSONInfo
	if encoding == ResponseEncodingJSONParsed {
		err := c.client.Call(&result, "getConfirmedTransaction", publicKey, encoding)
		if err != nil {
			return nil, fmt.Errorf("GetConfirmedTransaction with error %s", err.Error())
		}
		return &result, nil
	}

	if encoding == ResponseEncodingJSON {
		err := c.client.Call(&jresult, "getConfirmedTransaction", publicKey, encoding)
		if err != nil {
			return nil, fmt.Errorf("GetConfirmedTransaction with error %s", err.Error())
		}
		return &jresult, nil
	}

	if encoding == ResponseEncodingBase58 || encoding == ResponseEncodingBase64 {
		err := c.client.Call(&eresult, "getConfirmedTransaction", publicKey, encoding)
		if err != nil {
			return nil, fmt.Errorf("GetConfirmedTransaction with error %s", err.Error())
		}
		return &eresult, nil
	}

	return nil, fmt.Errorf("GetConfirmedTransaction with unsupport encoding")
}

// FeeCalculatorForBlockhashInfo is response for getFeeCalculatorForBlockhash
type FeeCalculatorForBlockhashInfo struct {
	FeeCalculator struct {
		LamportsPerSignature int `json:"lamportsPerSignature"`
	} `json:"feeCalculator"`
}

// GetFeeCalculatorForBlockhash Returns the fee calculator associated with
// the query blockhash, or null if the blockhash has expired
func (c *Connection) GetFeeCalculatorForBlockhash(publicKey string, commitment CommitmentValue) (
	*FeeCalculatorForBlockhashInfo, error) {
	var result RPCResult
	cmm := CommitmentConfig{
		Commitment: string(commitment),
	}
	var value FeeCalculatorForBlockhashInfo
	result.Value = &value
	err := c.client.Call(&result, "getFeeCalculatorForBlockhash", publicKey, cmm)
	if err != nil {
		return nil, fmt.Errorf("GetFeeCalculatorForBlockhash with error %s", err.Error())
	}
	if result.Value == nil {
		return nil, nil
	}
	return &value, nil
}

// FeeRateGovernorInfo is response for getFeeRateGovernor
type FeeRateGovernorInfo struct {
	FeeRateGovernor struct {
		BurnPercent                int `json:"burnPercent"`
		MaxLamportsPerSignature    int `json:"maxLamportsPerSignature"`
		MinLamportsPerSignature    int `json:"minLamportsPerSignature"`
		TargetLamportsPerSignature int `json:"targetLamportsPerSignature"`
		TargetSignaturesPerSlot    int `json:"targetSignaturesPerSlot"`
	} `json:"feeRateGovernor"`
}

// GetFeeRateGovernor Returns the fee rate governor information from the root bank
func (c *Connection) GetFeeRateGovernor() (*FeeRateGovernorInfo, error) {
	var result RPCResult
	var value FeeRateGovernorInfo
	result.Value = &value
	err := c.client.Call(&result, "getFeeRateGovernor")
	if err != nil {
		return nil, fmt.Errorf("GetFeeRateGovernor with error %s", err.Error())
	}
	return &value, nil
}

// FeesInfo is response of getFees
type FeesInfo struct {
	Blockhash     string `json:"blockhash"`
	FeeCalculator struct {
		LamportsPerSignature int `json:"lamportsPerSignature"`
	} `json:"feeCalculator"`
	LastValidSlot int `json:"lastValidSlot"`
}

// GetFees Returns a recent block hash from the ledger, a fee schedule that
// can be used to compute the cost of submitting a transaction using it,
// and the last slot in which the blockhash will be valid.
func (c *Connection) GetFees(commitment CommitmentValue) (*FeesInfo, error) {
	var result RPCResult
	var value FeesInfo
	cmm := CommitmentConfig{
		Commitment: string(commitment),
	}
	result.Value = &value
	err := c.client.Call(&result, "getFees", cmm)
	if err != nil {
		return nil, fmt.Errorf("GetFees with error %s", err.Error())
	}
	return &value, nil
}

// GetFirstAvailableBlock Returns the slot of the lowest confirmed block that
// has not been purged from the ledger
func (c *Connection) GetFirstAvailableBlock() (uint64, error) {
	var value uint64
	err := c.client.Call(&value, "getFirstAvailableBlock")
	if err != nil {
		return 0, fmt.Errorf("GetFirstAvailableBlock with error %s", err.Error())
	}
	return value, nil
}

// GetGenesisHash Returns the genesis hash
func (c *Connection) GetGenesisHash() (string, error) {
	var value string
	err := c.client.Call(&value, "getGenesisHash")
	if err != nil {
		return "", fmt.Errorf("GetGenesisHash with error %s", err.Error())
	}
	return value, nil
}

// IdentityInfo is response for getIdentity
type IdentityInfo struct {
	Identity string `json:"identity"`
}

// GetIdentity Returns the identity pubkey for the current node
func (c *Connection) GetIdentity() (string, error) {
	var result IdentityInfo
	err := c.client.Call(&result, "getIdentity")
	if err != nil {
		return "", fmt.Errorf("GetIdentity with error %s", err.Error())
	}
	return result.Identity, nil
}

// InflationGovernorInfo is response for getInflationGovernor
type InflationGovernorInfo struct {
	Foundation     float64 `json:"foundation"`
	FoundationTerm float64 `json:"foundationTerm"`
	Initial        float64 `json:"initial"`
	Taper          float64 `json:"taper"`
	Terminal       float64 `json:"terminal"`
}

// GetInflationGovernor Returns the current inflation governor
func (c *Connection) GetInflationGovernor(commitment CommitmentValue) (
	*InflationGovernorInfo, error) {
	var result InflationGovernorInfo
	cmm := CommitmentConfig{
		Commitment: string(commitment),
	}
	err := c.client.Call(&result, "getInflationGovernor", cmm)
	if err != nil {
		return nil, fmt.Errorf("GetInflationGovernor with error %s", err.Error())
	}
	return &result, nil
}

// InflationRateInfo is response for getInflationRate
type InflationRateInfo struct {
	Epoch      float64 `json:"epoch"`
	Foundation float64 `json:"foundation"`
	Total      float64 `json:"total"`
	Validator  float64 `json:"validator"`
}

// GetInflationRate Returns the specific inflation values for the current epoch
func (c *Connection) GetInflationRate() (*InflationRateInfo, error) {
	var result InflationRateInfo
	err := c.client.Call(&result, "getInflationRate")
	if err != nil {
		return nil, fmt.Errorf("GetInflationRate with error %s", err.Error())
	}
	return &result, nil
}

// GetLargestAccountsOpts is optinal for getLargestAccounts
type GetLargestAccountsOpts struct {
	Commitment string `json:"commitment"`
	Filter     string `json:"filter"`
}

// LargestAccountInfo is response for getLargestAccounts
type LargestAccountInfo struct {
	Lamports int    `json:"lamports"`
	Address  string `json:"address"`
}

// GetLargestAccounts Returns the 20 largest accounts, by lamport balance
func (c *Connection) GetLargestAccounts(commitment CommitmentValue,
	accountType AccountType) ([]LargestAccountInfo, error) {

	var value []LargestAccountInfo
	var result RPCResult
	result.Value = &value

	opts := GetLargestAccountsOpts{
		Commitment: string(commitment),
		Filter:     string(accountType),
	}
	err := c.client.Call(&result, "getLargestAccounts", opts)
	if err != nil {
		return nil, fmt.Errorf("GetLargestAccounts with error %s", err.Error())
	}
	return value, nil
}

// LeaderScheduleInfo is response for getLeaderSchedule
type LeaderScheduleInfo map[string][]uint64

// GetLeaderSchedule Returns the leader schedule for an epoch
func (c *Connection) GetLeaderSchedule(solt uint64, commitment CommitmentValue) (
	*LeaderScheduleInfo, error) {
	var result LeaderScheduleInfo = make(map[string][]uint64)
	cmm := CommitmentConfig{
		Commitment: string(commitment),
	}
	err := c.client.Call(&result, "getLeaderSchedule", solt, cmm)
	if err != nil {
		return nil, fmt.Errorf("GetLeaderSchedule with error %s", err.Error())
	}
	return &result, nil
}

// GetMinimumBalanceForRentExemption Returns minimum balance required to make account rent exempt.
func (c *Connection) GetMinimumBalanceForRentExemption(lenth uint64, commitment CommitmentValue) (
	uint64, error) {
	var result uint64
	cmm := CommitmentConfig{
		Commitment: string(commitment),
	}

	err := c.client.Call(&result, "getMinimumBalanceForRentExemption", lenth, cmm)
	if err != nil {
		return 0, fmt.Errorf("GetMinimumBalanceForRentExemption with error %s", err.Error())
	}
	return result, nil
}

// GetMultipleAccountsOpts is optinal for getMultipleAccounts
type GetMultipleAccountsOpts struct {
	Commitment struct {
		Commitment string `json:"commitment"`
	} `json:"commitment,omitempty"`
	Encoding  string `json:"encoding"`
	DataSlice struct {
		Offset int `json:"offset"`
		Length int `json:"length"`
	} `json:"dataSlice,omitempty"`
}

// MultipleAccountInfo is response for getMultipleAccounts
type MultipleAccountInfo struct {
	Data       []string `json:"data"`
	Executable bool     `json:"executable"`
	Lamports   int      `json:"lamports"`
	Owner      string   `json:"owner"`
	RentEpoch  int      `json:"rentEpoch"`
}

// GetMultipleAccounts Returns the account information for a list of Pubkeys
func (c *Connection) GetMultipleAccounts(publicKeys []string, opts GetMultipleAccountsOpts) (
	[]MultipleAccountInfo, error) {
	var value []MultipleAccountInfo
	var result RPCResult
	result.Value = &value

	err := c.client.Call(&result, "getMultipleAccounts", publicKeys, opts)
	if err != nil {
		return nil, fmt.Errorf("GetMultipleAccounts with error %s", err.Error())
	}
	return value, nil
}

// RecentBlockhashInfo response for getRecentBlockhash
type RecentBlockhashInfo struct {
	Blockhash     string `json:"blockhash"`
	FeeCalculator struct {
		LamportsPerSignature int `json:"lamportsPerSignature"`
	} `json:"feeCalculator"`
}

// GetRecentBlockhash Returns a recent block hash from the ledger,
// and a fee schedule that can be used to compute the cost of submitting a transaction using it
func (c *Connection) GetRecentBlockhash(commitment CommitmentValue) (
	*RecentBlockhashInfo, error) {
	var result RPCResult
	var value RecentBlockhashInfo
	cmm := CommitmentConfig{
		Commitment: string(commitment),
	}
	result.Value = &value
	err := c.client.Call(&result, "getRecentBlockhash", cmm)
	if err != nil {
		return nil, fmt.Errorf("GetRecentBlockhash with error %s", err.Error())
	}
	return &value, nil
}

// RecentPerformanceSampleInfo is response for getRecentPerformanceSamples
type RecentPerformanceSampleInfo struct {
	NumSlots         int `json:"numSlots"`
	NumTransactions  int `json:"numTransactions"`
	SamplePeriodSecs int `json:"samplePeriodSecs"`
	Slot             int `json:"slot"`
}

// GetRecentPerformanceSamples Returns a list of recent performance samples,
// in reverse slot order. Performance samples are taken every 60 seconds
// and include the number of transactions and slots that occur in a given time window.
func (c *Connection) GetRecentPerformanceSamples(limit uint64) (
	[]RecentPerformanceSampleInfo, error) {
	var value []RecentPerformanceSampleInfo
	err := c.client.Call(&value, "getRecentPerformanceSamples", limit)
	if err != nil {
		return nil, fmt.Errorf("GetRecentPerformanceSamples with error %s", err.Error())
	}
	return value, nil
}

// GetSignatureStatusesOpts is optional for GetSignatureStatuses
type GetSignatureStatusesOpts struct {
	SearchTransactionHistory bool `json:"searchTransactionHistory"`
}

// SignatureStatuseInfo is response for GetSignatureStatuses
type SignatureStatuseInfo struct {
	Slot          int         `json:"slot"`
	Confirmations interface{} `json:"confirmations"`
	Err           interface{} `json:"err"`
	Status        struct {
		Ok interface{} `json:"Ok"`
	} `json:"status"`
}

// GetSignatureStatuses Returns the statuses of a list of signatures. Unless
//the searchTransactionHistory configuration parameter is included, this method
//only searches the recent status cache of signatures, which retains statuses
//for all active slots plus MAX_RECENT_BLOCKHASHES rooted slots.
func (c *Connection) GetSignatureStatuses(signatures []string, searchTransactionHistory bool) (
	[]SignatureStatuseInfo, error) {
	opts := GetSignatureStatusesOpts{
		SearchTransactionHistory: searchTransactionHistory,
	}
	var value []SignatureStatuseInfo
	err := c.client.Call(&value, "getSignatureStatuses", signatures, opts)
	if err != nil {
		return nil, fmt.Errorf("GetSignatureStatuses with error %s", err.Error())
	}
	return value, nil
}

// GetSlot Returns the current slot the node is processing
func (c *Connection) GetSlot(commitment CommitmentValue) (uint64, error) {
	var result uint64
	cmm := CommitmentConfig{
		Commitment: string(commitment),
	}
	err := c.client.Call(&result, "getSlot", cmm)
	if err != nil {
		return 0, fmt.Errorf("MinimumLedgerSlot with error %s", err.Error())
	}
	return result, nil
}

// GetSlotLeader Returns the current slot leader
func (c *Connection) GetSlotLeader(commitment CommitmentValue) (string, error) {
	var result string
	cmm := CommitmentConfig{
		Commitment: string(commitment),
	}
	err := c.client.Call(&result, "getSlotLeader", cmm)
	if err != nil {
		return "", fmt.Errorf("GetSlotLeader with error %s", err.Error())
	}
	return result, nil
}

// GetStakeActivationOpts is optional for GetStakeActivation
type GetStakeActivationOpts struct {
	Commitment struct {
		Commitment string `json:"commitment"`
	} `json:"commitment,omitempty"`
	Epoch int `json:"epoch,omitempty"`
}

// StakeActivationInfo is response for GetStakeActivation
type StakeActivationInfo struct {
	Active   int    `json:"active"`
	Inactive int    `json:"inactive"`
	State    string `json:"state"`
}

// GetStakeActivation Returns epoch activation information for a stake account
func (c *Connection) GetStakeActivation(publicKey string, opts GetStakeActivationOpts) (
	*StakeActivationInfo, error) {
	var result StakeActivationInfo
	err := c.client.Call(&result, "getStakeActivation", publicKey, opts)
	if err != nil {
		return nil, fmt.Errorf("GetSlotLeader with error %s", err.Error())
	}
	return &result, nil
}

//SupplyInfo is response for GetSupply
type SupplyInfo struct {
	Circulating            int      `json:"circulating"`
	NonCirculating         int      `json:"nonCirculating"`
	NonCirculatingAccounts []string `json:"nonCirculatingAccounts"`
	Total                  int      `json:"total"`
}

// GetSupply Returns information about the current supply.
func (c *Connection) GetSupply(commitment CommitmentValue) (*SupplyInfo, error) {
	var value SupplyInfo
	var result RPCResult
	cmm := CommitmentConfig{
		Commitment: string(commitment),
	}
	result.Value = &value
	err := c.client.Call(&result, "getSupply", cmm)
	if err != nil {
		return nil, fmt.Errorf("GetSlotLeader with error %s", err.Error())
	}
	return &value, nil
}

// TokenAccountBalanceInfo is response for GetTokenAccountBalance
type TokenAccountBalanceInfo struct {
	UIAmount float64 `json:"uiAmount"`
	Amount   string  `json:"amount"`
	Decimals int     `json:"decimals"`
}

//GetTokenAccountBalance Returns the token balance of an SPL Token account
func (c *Connection) GetTokenAccountBalance(publicKey string, commitment CommitmentValue) (
	*TokenAccountBalanceInfo, error) {
	var value TokenAccountBalanceInfo
	var result RPCResult
	cmm := CommitmentConfig{
		Commitment: string(commitment),
	}
	result.Value = &value
	err := c.client.Call(&result, "getTokenAccountBalance", publicKey, cmm)
	if err != nil {
		return nil, fmt.Errorf("GetTokenAccountBalance with error %s", err.Error())
	}
	return &value, nil
}

/***************************/

// MinimumLedgerSlot Returns the lowest slot that the node has information
//about in its ledger. This value may increase over time if the node is
//configured to purge older ledger data
func (c *Connection) MinimumLedgerSlot() (uint64, error) {
	var result uint64
	err := c.client.Call(&result, "minimumLedgerSlot")
	if err != nil {
		return 0, fmt.Errorf("MinimumLedgerSlot with error %s", err.Error())
	}
	return result, nil
}

// RequestAirdrop Requests an airdrop of lamports to a Pubkey
func (c *Connection) RequestAirdrop(publicKey string, lamports uint64, commitment CommitmentValue) (
	string, error) {

	var result string
	cmm := CommitmentConfig{
		Commitment: string(commitment),
	}
	err := c.client.Call(&result, "requestAirdrop", publicKey, lamports, cmm)
	if err != nil {
		return "", fmt.Errorf("RequestAirdrop with error %s", err.Error())
	}
	return result, nil
}

// SetLogFilter Sets the log filter on the validator
func (c *Connection) SetLogFilter(filter string) error {
	var result interface{}
	err := c.client.Call(&result, "setLogFilter", filter)
	if err != nil {
		return fmt.Errorf("SetLogFilter with error %s", err.Error())
	}
	return nil
}

// ValidatorExit If a validator boots with RPC exit enabled
// (--enable-rpc-exit parameter), this request causes the validator to exit.
func (c *Connection) ValidatorExit() (bool, error) {
	var result bool
	err := c.client.Call(&result, "validatorExit")
	if err != nil {
		return false, fmt.Errorf("SetLogFilter with error %s", err.Error())
	}
	return result, nil
}
