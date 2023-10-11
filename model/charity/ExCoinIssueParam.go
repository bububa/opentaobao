package charity

// ExCoinIssueParam 结构体
type ExCoinIssueParam struct {
	// 开放标识
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 业务编码
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 幂等ID
	IdempotentId string `json:"idempotent_id,omitempty" xml:"idempotent_id,omitempty"`
	// 渠道编码
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 爱豆发放时间
	IssueTime int64 `json:"issue_time,omitempty" xml:"issue_time,omitempty"`
	// 行为ID
	ActionId int64 `json:"action_id,omitempty" xml:"action_id,omitempty"`
}
