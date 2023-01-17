package idleisv

// IdleDepositTopSubVo 结构体
type IdleDepositTopSubVo struct {
	// 保证金余额是否足够
	BalanceEnough bool `json:"balance_enough,omitempty" xml:"balance_enough,omitempty"`
}
