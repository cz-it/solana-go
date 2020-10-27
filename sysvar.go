package solana

// Sysvar Programs
var (
	SysvarClock, _             = NewPublicKey("SysvarC1ock11111111111111111111111111111111")
	SysvarRecentBlockHashes, _ = NewPublicKey("SysvarRecentB1ockHashes11111111111111111111")
	SysvarRent, _              = NewPublicKey("SysvarRent111111111111111111111111111111111")
	SysvarRewards, _           = NewPublicKey("SysvarRewards111111111111111111111111111111")
	SysvarStakeHistory, _      = NewPublicKey("SysvarStakeHistory1111111111111111111111111")
)
