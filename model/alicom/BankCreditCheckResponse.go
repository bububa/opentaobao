package alicom

// BankCreditCheckResponse 结构体
type BankCreditCheckResponse struct {
	// 能否申请
	CanApply bool `json:"can_apply,omitempty" xml:"can_apply,omitempty"`
}
