package solana

import (
	"fmt"
)

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
func (c *Connection) GetEpochInfo(commitment CommitmentConfig) (*EpochInfo, error) {
	var result EpochInfo
	err := c.client.Call(&result, "getEpochInfo", commitment)
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
	// DataSlice  struct {
	// 	Offset int `json:"offset"`
	// 	Length int `json:"length"`
	// } `json:"dataSlice"`
}

// AccountInfo is response of getAccountInfo
type AccountInfo struct {
	Context struct {
		Slot int `json:"slot"`
	} `json:"context"`
	Value struct {
		Data       []string `json:"data"`
		Executable bool     `json:"executable"`
		Lamports   int64    `json:"lamports"`
		Owner      string   `json:"owner"`
		RentEpoch  int      `json:"rentEpoch"`
	} `json:"value"`
}

// GetAccountInfo Fetch all the account info for the specified public key, return with context
func (c *Connection) GetAccountInfo(publicKey string, opts *GetAccountInfoOpts) (*AccountInfo, error) {
	var result AccountInfo
	err := c.client.Call(&result, "getAccountInfo", publicKey, opts)
	if err != nil {
		return nil, fmt.Errorf("GetAccountInfo with error %s", err.Error())
	}
	return &result, nil
}
