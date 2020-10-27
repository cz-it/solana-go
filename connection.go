package solana

import (
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"
)

// Connection is a connection to API node
type Connection struct {
	client *rpc.Client
}

// NewConnection create a connection to endpoint
func NewConnection(endpoint, commitment string) *Connection {
	c := &Connection{}
	c.client, _ = rpc.Dial(endpoint)
	return c
}

// EpochInfo is response of getEpochInfo
type EpochInfo struct {
	AbsoluteSlot int `json:"absoluteSlot"`
	BlockHeight  int `json:"blockHeight"`
	Epoch        int `json:"epoch"`
	SlotIndex    int `json:"slotIndex"`
	SlotsInEpoch int `json:"slotsInEpoch"`
}

// CommitmentConfig is commitment
type CommitmentConfig struct {
	Commitment string `json:"commitment"`
}

// GetEpochInfo Fetch the Epoch Info parameters
func (c *Connection) GetEpochInfo(commitment CommitmentConfig) (epochInfo *EpochInfo, err error) {
	var result EpochInfo
	err = c.client.Call(&result, "getEpochInfo", commitment)
	if err != nil {
		return nil, fmt.Errorf("GetEpochInfo with error %s", err.Error())
	}
	return &result, nil
}
