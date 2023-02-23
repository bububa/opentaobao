package train

// ChangeIssueRq 结构体
type ChangeIssueRq struct {
	// 改签票信息列表
	ChangeTickets []ChangeTicketDto `json:"change_tickets,omitempty" xml:"change_tickets>change_ticket_dto,omitempty"`
	// 改签票出发站
	ChangeFromStation string `json:"change_from_station,omitempty" xml:"change_from_station,omitempty"`
	// 支付宝交易流水号
	AlipayTradeNo string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// 支付宝账号
	AlipayAccount string `json:"alipay_account,omitempty" xml:"alipay_account,omitempty"`
	// 改签票到达站
	ChangeToStation string `json:"change_to_station,omitempty" xml:"change_to_station,omitempty"`
	// 改签票123N6订单号
	SequenceNo string `json:"sequence_no,omitempty" xml:"sequence_no,omitempty"`
	// 改签票发车时间
	ChangeFromDateTime string `json:"change_from_date_time,omitempty" xml:"change_from_date_time,omitempty"`
	// 改签票到达时间
	ChangeToDateTime string `json:"change_to_date_time,omitempty" xml:"change_to_date_time,omitempty"`
	// 改签票发车日期
	ChangeTrainDate string `json:"change_train_date,omitempty" xml:"change_train_date,omitempty"`
	// 改签票车次号
	ChangeTrainCode string `json:"change_train_code,omitempty" xml:"change_train_code,omitempty"`
	// 改签出票失败msg
	IssueFailMsg string `json:"issue_fail_msg,omitempty" xml:"issue_fail_msg,omitempty"`
	// 改签失败出票回填code
	IssueFailCode int64 `json:"issue_fail_code,omitempty" xml:"issue_fail_code,omitempty"`
	// 代理商id
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 改签申请单id
	ChangeApplyId int64 `json:"change_apply_id,omitempty" xml:"change_apply_id,omitempty"`
	// 改签结算方式 0现金 1电子
	ChangeSettlementMode int64 `json:"change_settlement_mode,omitempty" xml:"change_settlement_mode,omitempty"`
	// 淘宝主单单号
	TpOrderId int64 `json:"tp_order_id,omitempty" xml:"tp_order_id,omitempty"`
	// ttp单号
	TtpOrderId int64 `json:"ttp_order_id,omitempty" xml:"ttp_order_id,omitempty"`
	// 出票状态 1成功 2失败
	IssueStatus int64 `json:"issue_status,omitempty" xml:"issue_status,omitempty"`
}
