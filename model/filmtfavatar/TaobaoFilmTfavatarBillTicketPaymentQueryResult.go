package filmtfavatar

// TaobaoFilmTfavatarBillTicketPaymentQueryResult 结构体
type TaobaoFilmTfavatarBillTicketPaymentQueryResult struct {
	// 返回码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 请求id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msg
	ReturnMessage string `json:"return_message,omitempty" xml:"return_message,omitempty"`
	// 返回参数
	ReturnValue *ReturnValue `json:"return_value,omitempty" xml:"return_value,omitempty"`
}
