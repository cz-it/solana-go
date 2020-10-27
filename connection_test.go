package solana

import (
	"testing"
)

const MainNet string = "https://api.mainnet-beta.solana.com"

func TestGetEpochInfo(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	cmm := CommitmentConfig{Commitment: "recent"}
	c.GetEpochInfo(cmm)
}
