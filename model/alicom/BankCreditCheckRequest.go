package alicom

// BankCreditCheckRequest 结构体
type BankCreditCheckRequest struct {
	// 银行的编码
	BankCode string `json:"bank_code,omitempty" xml:"bank_code,omitempty"`
	// 用户的openId
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
}
