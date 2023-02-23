package train

// QueryChangeDetailRq 结构体
type QueryChangeDetailRq struct {
	// 代理商id
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 改签申请单id
	ChangeApplyId int64 `json:"change_apply_id,omitempty" xml:"change_apply_id,omitempty"`
	// 淘宝主单id
	TpOrderId int64 `json:"tp_order_id,omitempty" xml:"tp_order_id,omitempty"`
	// ttp单号
	TtpOrderId int64 `json:"ttp_order_id,omitempty" xml:"ttp_order_id,omitempty"`
}
