package alsc

// PointFlowOpenInfo 结构体
type PointFlowOpenInfo struct {
	// 业务场景
	BizType string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 业务类型描述
	BizTypeDescription string `json:"biz_type_description,omitempty" xml:"biz_type_description,omitempty"`
	// 变更积分
	ChangePoint int64 `json:"change_point,omitempty" xml:"change_point,omitempty"`
	// 流水id
	FlowId string `json:"flow_id,omitempty" xml:"flow_id,omitempty"`
	// 操作员名
	OperatorName string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
	// 交易单号
	OutBizId string `json:"out_biz_id,omitempty" xml:"out_biz_id,omitempty"`
	// 变更原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 剩余积分
	RemainPoint int64 `json:"remain_point,omitempty" xml:"remain_point,omitempty"`
	// 店铺名称
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// 交易时间
	Time string `json:"time,omitempty" xml:"time,omitempty"`
}
