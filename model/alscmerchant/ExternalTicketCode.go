package alscmerchant

// ExternalTicketCode 结构体
type ExternalTicketCode struct {
	// 码的可核销份数
	Num string `json:"num,omitempty" xml:"num,omitempty"`
	// 凭证码值，必填
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 本地生活凭证id(奇门发码通知中的ticketId)，必填
	TicketId string `json:"ticket_id,omitempty" xml:"ticket_id,omitempty"`
	// 支持核销的三方url
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}
