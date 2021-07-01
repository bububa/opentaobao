package aliospay

// GetTokenResponse 结构体
type GetTokenResponse struct {
	// 支付token
	PayToken string `json:"pay_token,omitempty" xml:"pay_token,omitempty"`
}
