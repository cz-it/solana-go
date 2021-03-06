package solana

import (
	"testing"
)

const MainNet string = "https://api.mainnet-beta.solana.com"
const TestPublicKey = "HSfwVfB7RUF1SKCd4yrz8KZp7TU262Y5BeZZN1tdCTVk"
const TestSig = "4zcDcRS5LFzvCzgZbZWofQD8wbi5DpMQUbjrWAQEpistzTiwXaJLvVoUtnaXLzxPvXwM7mvfhHXHG7kVZ85NUAJc"
const TokenProgram = "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"
const TestSPLAccount = "6xe4WTMLrcrxae196V4PLhNpqWgzmFqJw1CLVvWHkw4U"

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

func TestGetFeeCalculatorForBlockhash(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.GetFeeCalculatorForBlockhash(TokenProgram, CommitmentRecent)
	if err != nil {
		t.Errorf("get fee calculator info error:%s", err.Error())
	}
	t.Logf("fee: %v\n", info)
}

func TestGetFeeRateGovernor(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.GetFeeRateGovernor()
	if err != nil {
		t.Errorf("get fee rate info error:%s", err.Error())
	}
	t.Logf("fee rate: %v\n", info)
}

func TestGetFees(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.GetFees(CommitmentRecent)
	if err != nil {
		t.Errorf("get fees info error:%s", err.Error())
	}
	t.Logf("fees: %v\n", info)
}

func TestGetFirstAvailableBlock(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.GetFirstAvailableBlock()
	if err != nil {
		t.Errorf("get first block info error:%s", err.Error())
	}
	t.Logf("block : %v\n", info)
}

func TestGetGenesisHash(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.GetGenesisHash()
	if err != nil {
		t.Errorf("get genesis block info error:%s", err.Error())
	}
	t.Logf("block : %v\n", info)
}

func TestGetIdentity(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.GetIdentity()
	if err != nil {
		t.Errorf("get  identity info error:%s", err.Error())
	}
	t.Logf("identity : %v\n", info)
}

func TestGetInflationGovernor(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.GetInflationGovernor(CommitmentRecent)
	if err != nil {
		t.Errorf("get InflationGovernor info error:%s", err.Error())
	}
	t.Logf("Inflation: %v\n", info)
}

func TestGetInflationRate(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.GetInflationRate()
	if err != nil {
		t.Errorf("get GetInflationRate info error:%s", err.Error())
	}
	t.Logf("InflationRate: %v\n", info)
}

func TestGetLargestAccounts(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.GetLargestAccounts(CommitmentRecent, AccountTypeCirculating)
	if err != nil {
		t.Errorf("GetLargestAccounts info error:%s", err.Error())
	}
	t.Logf("GetLargestAccounts: %v\n", info)
}

func TestGetLeaderSchedule(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.GetLeaderSchedule(104, CommitmentRecent)
	if err != nil {
		t.Errorf("GetLeaderSchedule info error:%s", err.Error())
	}
	t.Logf("GetLeaderSchedule: %v\n", info)
}

func TestSetLogFilter(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	err := c.SetLogFilter("solana_core=debug")
	if err != nil {
		t.Errorf("SetLogFilter info error:%s", err.Error())
	}
	t.Logf("SetLogFilter success \n")
}

func TestValidatorExit(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.ValidatorExit()
	if err != nil {
		t.Errorf("SetLogFilter info error:%s", err.Error())
	}
	t.Logf("ValidatorExit info:%v\n", info)
}

func TestRequestAirdrop(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.RequestAirdrop(TestPublicKey, 100, CommitmentRecent)
	if err != nil {
		//t.Errorf("RequestAirdrop info error:%s", err.Error())
	}
	t.Logf("RequestAirdrop info:%v\n", info)
}

func TestMinimumLedgerSlot(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.MinimumLedgerSlot()
	if err != nil {
		t.Errorf("MinimumLedgerSlot info error:%s", err.Error())
	}
	t.Logf("MinimumLedgerSlot info:%v\n", info)
}

func TestGetMinimumBalanceForRentExemption(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.GetMinimumBalanceForRentExemption(50, CommitmentRecent)
	if err != nil {
		t.Errorf("GetMinimumBalanceForRentExemption info error:%s", err.Error())
	}
	t.Logf("GetMinimumBalanceForRentExemption info:%v\n", info)
}

func TestGetMultipleAccounts(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	opts := GetMultipleAccountsOpts{
		Encoding: string(ResponseEncodingBase58),
		DataSlice: struct {
			Offset int `json:"offset"`
			Length int `json:"length"`
		}{
			Offset: 0,
			Length: 10,
		},
	}
	publicKeys := []string{TestPublicKey}
	info, err := c.GetMultipleAccounts(publicKeys, opts)
	if err != nil {
		t.Errorf("GetMultipleAccounts info error:%s", err.Error())
	}
	t.Logf("GetMultipleAccounts info:%v\n", info)
}

func TestGetRecentBlockhash(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.GetRecentBlockhash(CommitmentRecent)
	if err != nil {
		t.Errorf("GetRecentBlockhash info error:%s", err.Error())
	}
	t.Logf("GetRecentBlockhash info:%v\n", info)
}

func TestGetRecentPerformanceSamples(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.GetRecentPerformanceSamples(10)
	if err != nil {
		t.Errorf("GetRecentPerformanceSamples info error:%s", err.Error())
	}
	t.Logf("GetRecentPerformanceSamples info:%v\n", info)
}

func TestGetSignatureStatuses(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.GetSignatureStatuses([]string{TestSig}, true)
	if err != nil {
		//t.Errorf("GetSignatureStatuses info error:%s", err.Error())
	}
	t.Logf("GetSignatureStatuses info:%v\n", info)
}

func TestGetSlot(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.GetSlot(CommitmentRecent)
	if err != nil {
		t.Errorf("GetSlot info error:%s", err.Error())
	}
	t.Logf("GetSlot info:%v\n", info)
}

func TestGetSlotLeader(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.GetSlotLeader(CommitmentRecent)
	if err != nil {
		t.Errorf("GetSlotLeader info error:%s", err.Error())
	}
	t.Logf("GetSlotLeader info:%v\n", info)
}

func TestGetStakeActivation(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	opts := GetStakeActivationOpts{
		Epoch: 4,
	}
	info, err := c.GetStakeActivation(TestPublicKey, opts)
	if err != nil {
		//t.Errorf("GetStakeActivation info error:%s", err.Error())
	}
	t.Logf("GetStakeActivation info:%v\n", info)
}

func TestGetSupply(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.GetSupply(CommitmentRecent)
	if err != nil {
		t.Errorf("GetSupply info error:%s", err.Error())
	}
	t.Logf("GetSupply info:%v\n", info)
}

func TestGetTokenAccountBalance(t *testing.T) {
	c := NewConnection(MainNet, "recent")
	info, err := c.GetTokenAccountBalance(TestSPLAccount, CommitmentRecent)
	if err != nil {
		t.Errorf("GetTokenAccountBalance info error:%s", err.Error())
	}
	t.Logf("GetTokenAccountBalance info:%v\n", info)
}
