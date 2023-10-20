package filmtfavatar

// TaobaoFilmTfavatarBillSalePaymentQueryViiResult 结构体
type TaobaoFilmTfavatarBillSalePaymentQueryViiResult struct {
	// 错误信息
	ReturnMessage string `json:"return_message,omitempty" xml:"return_message,omitempty"`
	// 错误码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 请求ID(排查问题用)
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	ReturnValue *ReturnValue `json:"return_value,omitempty" xml:"return_value,omitempty"`
}
