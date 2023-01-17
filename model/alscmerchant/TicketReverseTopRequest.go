package alscmerchant

// TicketReverseTopRequest 结构体
type TicketReverseTopRequest struct {
	// 凭证核销流水id
	TicketUseTransId string `json:"ticket_use_trans_id,omitempty" xml:"ticket_use_trans_id,omitempty"`
	// 业务发生时间，为空取调用时间
	BizDate string `json:"biz_date,omitempty" xml:"biz_date,omitempty"`
	// 请求id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 口碑凭证id
	TicketId string `json:"ticket_id,omitempty" xml:"ticket_id,omitempty"`
}
