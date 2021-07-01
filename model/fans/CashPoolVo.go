package fans

// CashPoolVo 结构体
type CashPoolVo struct {
	// 付款url
	PayUrl string `json:"pay_url,omitempty" xml:"pay_url,omitempty"`
	// 奖金池id
	CashPoolId int64 `json:"cash_pool_id,omitempty" xml:"cash_pool_id,omitempty"`
}
