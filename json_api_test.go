package solana

import (
	"testing"
)

const MainNet string = "https://api.mainnet-beta.solana.com"
const TestPublicKey = "HSfwVfB7RUF1SKCd4yrz8KZp7TU262Y5BeZZN1tdCTVk"
const TestSig = "4zcDcRS5LFzvCzgZbZWofQD8wbi5DpMQUbjrWAQEpistzTiwXaJLvVoUtnaXLzxPvXwM7mvfhHXHG7kVZ85NUAJc"

func TestGetEpochInfo(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.GetEpochInfo(CommitmentRecent)
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

func TestGetBalance(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.GetBalance(TestPublicKey, CommitmentRecent)
	if err != nil {
		t.Errorf("get balance info error:%s", err.Error())
	}
	t.Logf("balance : %v\n", info)
}

func TestGetBlockCommitment(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.GetBlockCommitment(4)
	if err != nil {
		t.Errorf("get block commitment info error:%s", err.Error())
	}
	t.Logf("commitment: %v\n", info)
}

func TestGetBlockTime(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.GetBlockCommitment(44902394)
	if err != nil {
		t.Errorf("get block time error:%s", err.Error())
	}
	t.Logf("timestamp: %v\n", info)
}

func TestGetClusterNodes(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.GetClusterNodes()
	if err != nil {
		t.Errorf("get nodes info error:%s", err.Error())
	}
	t.Logf("nodes: %v\n", info)
}

func TestGetConfirmedBlock(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.GetConfirmedBlock(44902219, "jsonParsed")
	if err != nil {
		t.Errorf("get confirmed info error:%s", err.Error())
	}
	rst := info.(*ConfirmedBlockInfo)
	t.Logf("block info: %v\n", rst)

	info2, err := c.GetConfirmedBlock(4, "base58")
	if err != nil {
		t.Errorf("get confirmed info error:%s", err.Error())
	}
	rst2 := info2.(*ConfirmedBlockEncodedInfo)
	t.Logf("block info: %v\n", rst2)

	info3, err := c.GetConfirmedBlock(4, "json")
	if err != nil {
		t.Errorf("get confirmed info error:%s", err.Error())
	}
	rst3 := info3.(*ConfirmedBlockJSONInfo)
	t.Logf("block info: %v\n", rst3)
}

func TestGetConfirmedBlocks(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.GetConfirmedBlocks(4, 8)
	if err != nil {
		t.Errorf("get confirmed blocks info error:%s", err.Error())
	}
	t.Logf("slots: %v\n", info)
}

func TestGetConfirmedBlocksWithLimit(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.GetConfirmedBlocksWithLimit(4, 3)
	if err != nil {
		t.Errorf("get confirmed blocks info error:%s", err.Error())
	}
	t.Logf("slots: %v\n", info)
}

func TestGetConfirmedSignaturesForAddress2(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	opts := GetConfirmedSignaturesForAddress2Opts{
		Limit: 5,
	}
	info, err := c.GetConfirmedSignaturesForAddress2(TestPublicKey, opts)
	if err != nil {
		t.Errorf("get confirmed Signatures info error:%s", err.Error())
	}
	t.Logf("Signatures: %v\n", info)
}

func TestGetConfirmedTransaction(t *testing.T) {

	c := NewConnection(MainNet, "recent")
	info, err := c.GetConfirmedTransaction(TestSig, "jsonParsed")
	if err != nil {
		t.Errorf("get confirmed info error:%s", err.Error())
	}
	rst := info.(*ConfirmedTransactionInfo)
	t.Logf("block info: %v\n", rst)

	info2, err := c.GetConfirmedTransaction(TestSig, "base58")
	if err != nil {
		t.Errorf("get confirmed info error:%s", err.Error())
	}
	rst2 := info2.(*ConfirmedTransactionEncodedInfo)
	t.Logf("block info: %v\n", rst2)

	info3, err := c.GetConfirmedTransaction(TestSig, "json")
	if err != nil {
		t.Errorf("get confirmed info error:%s", err.Error())
	}
	rst3 := info3.(*ConfirmedTransactionJSONInfo)
	t.Logf("block info: %v\n", rst3)
}
