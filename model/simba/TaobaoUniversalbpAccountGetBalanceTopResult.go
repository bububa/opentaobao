package simba

// TaobaoUniversalbpAccountGetBalanceTopResult 结构体
type TaobaoUniversalbpAccountGetBalanceTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	TopAccountBalanceVO *TopAccountBalanceVo `json:"top_account_balance_v_o,omitempty" xml:"top_account_balance_v_o,omitempty"`
}
