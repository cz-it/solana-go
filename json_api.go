package solana

import (
	"fmt"
)

// RPCResult  is result for JSON RPC
type RPCResult struct {
	Context struct {
		Slot int `json:"slot"`
	} `json:"context"`
	Value interface{} `josn:"value"`
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
