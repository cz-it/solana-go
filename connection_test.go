package solana

import (
	"testing"
)

const MainNet string = "https://api.mainnet-beta.solana.com"
const TestPublicKey = "HSfwVfB7RUF1SKCd4yrz8KZp7TU262Y5BeZZN1tdCTVk"

func TestGetEpochInfo(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	cmm := CommitmentConfig{Commitment: "recent"}
	info, err := c.GetEpochInfo(cmm)
	if err != nil {
		t.Errorf("get epoch info error:%s", err.Error())
	}
	t.Logf("info: %v\n", info)
}

func TestGetEpochSchedule(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.GetEpochSchedule()
	if err != nil {
		t.Errorf("get epoch schedule error:%s", err.Error())
	}
	t.Logf("schedule: %v\n", info)
}

func TestGetAccountInfo(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	opts := &GetAccountInfoOpts{
		Encoding:   "jsonParsed",
		Commitment: "rencent",
	}
	info, err := c.GetAccountInfo(TestPublicKey, opts)
	if err != nil {
		t.Errorf("get account info error:%s", err.Error())
	}
	t.Logf("account info: %v\n", info)
}
