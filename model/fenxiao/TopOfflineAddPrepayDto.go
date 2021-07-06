package fenxiao

// TopOfflineAddPrepayDto 结构体
type TopOfflineAddPrepayDto struct {
	// 收款人账号
	ReceiverAccountNum string `json:"receiver_account_num,omitempty" xml:"receiver_account_num,omitempty"`
	// 外部系统支付流水Id，用于商家上传流水时去重(外部保证唯一）
	OuterPayId string `json:"outer_pay_id,omitempty" xml:"outer_pay_id,omitempty"`
	// 销商淘宝nick
	DistNick string `json:"dist_nick,omitempty" xml:"dist_nick,omitempty"`
	// 出票人全称
	DrawerFullName string `json:"drawer_full_name,omitempty" xml:"drawer_full_name,omitempty"`
	// 付款行行号
	PayBankNum string `json:"pay_bank_num,omitempty" xml:"pay_bank_num,omitempty"`
	// 出票人账号
	DrawerAccountNum string `json:"drawer_account_num,omitempty" xml:"drawer_account_num,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 出票日期
	TicketIssueDate string `json:"ticket_issue_date,omitempty" xml:"ticket_issue_date,omitempty"`
	// 收款开户银行
	ReceiverBankFullName string `json:"receiver_bank_full_name,omitempty" xml:"receiver_bank_full_name,omitempty"`
	// 承兑票据号（承兑必填）
	TicketId string `json:"ticket_id,omitempty" xml:"ticket_id,omitempty"`
	// 汇票到期日期
	AcceptDate string `json:"accept_date,omitempty" xml:"accept_date,omitempty"`
	// 收款人全称
	ReceiverFullName string `json:"receiver_full_name,omitempty" xml:"receiver_full_name,omitempty"`
	// 付款行全称
	PayBankFullName string `json:"pay_bank_full_name,omitempty" xml:"pay_bank_full_name,omitempty"`
	// 资金流水类型：1.纸质承兑； 2.电子承兑；3.现金；4.优惠返点；5.奖励
	FlowType int64 `json:"flow_type,omitempty" xml:"flow_type,omitempty"`
	// 金额，单位分（必须为正数）
	TicketMoney int64 `json:"ticket_money,omitempty" xml:"ticket_money,omitempty"`
}
