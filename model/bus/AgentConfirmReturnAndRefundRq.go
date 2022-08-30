package bus

// AgentConfirmReturnAndRefundRq 结构体
type AgentConfirmReturnAndRefundRq struct {
	// 商家订单号
	AgentOrderId string `json:"agent_order_id,omitempty" xml:"agent_order_id,omitempty"`
	// 退款金额,单位元;退票成功时必填
	AgentRefundAmount string `json:"agent_refund_amount,omitempty" xml:"agent_refund_amount,omitempty"`
	// 退款资金号唯一ID
	AgentRefundTransId string `json:"agent_refund_trans_id,omitempty" xml:"agent_refund_trans_id,omitempty"`
	// 退票时间点
	AgentReturnTime string `json:"agent_return_time,omitempty" xml:"agent_return_time,omitempty"`
	// 按票的维度进行退票时，商家票号
	AgentTicketId string `json:"agent_ticket_id,omitempty" xml:"agent_ticket_id,omitempty"`
	// 发车时间
	DepartDate string `json:"depart_date,omitempty" xml:"depart_date,omitempty"`
	// 退票乘客证件号(按票退时必填)
	PassengerIdNum string `json:"passenger_id_num,omitempty" xml:"passenger_id_num,omitempty"`
	// 退票乘客姓名
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 退款乘客手机号
	PassengerPhone string `json:"passenger_phone,omitempty" xml:"passenger_phone,omitempty"`
	// 扩展属性json
	AgentExtAttr string `json:"agent_ext_attr,omitempty" xml:"agent_ext_attr,omitempty"`
	// offline:线下退票;online:线上退票
	AgentReturnMode string `json:"agent_return_mode,omitempty" xml:"agent_return_mode,omitempty"`
	// 退票场景，common:普通退款场景 force: 不可抗力
	RefundScene string `json:"refund_scene,omitempty" xml:"refund_scene,omitempty"`
	// 通知用户退票原因
	RefundSceneOfficialReasonCode string `json:"refund_scene_official_reason_code,omitempty" xml:"refund_scene_official_reason_code,omitempty"`
	// 商家退票状态;1-成功,2-失败
	AgentReturnTicketStatus int64 `json:"agent_return_ticket_status,omitempty" xml:"agent_return_ticket_status,omitempty"`
	// 商家退票类型;0-按票退,1-按单退
	AgentReturnTicketType int64 `json:"agent_return_ticket_type,omitempty" xml:"agent_return_ticket_type,omitempty"`
	// 平台单号
	MainBizOrderId int64 `json:"main_biz_order_id,omitempty" xml:"main_biz_order_id,omitempty"`
	// 退票结果编码,(退票失败时必填,参考标准错误码)
	AgentReturnTicketCode int64 `json:"agent_return_ticket_code,omitempty" xml:"agent_return_ticket_code,omitempty"`
}
