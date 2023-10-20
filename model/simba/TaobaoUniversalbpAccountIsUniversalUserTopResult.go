package simba

// TaobaoUniversalbpAccountIsUniversalUserTopResult 结构体
type TaobaoUniversalbpAccountIsUniversalUserTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	TopAccountStatusVO *TopAccountStatusVo `json:"top_account_status_v_o,omitempty" xml:"top_account_status_v_o,omitempty"`
}
