package wdk

// MemberQueryRequest 结构体
type MemberQueryRequest struct {
	// 商户号
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 扩展字段json字符串
	MemberAttributes string `json:"member_attributes,omitempty" xml:"member_attributes,omitempty"`
	// 会员类型
	AccountType string `json:"account_type,omitempty" xml:"account_type,omitempty"`
	// 会员号
	AccountId string `json:"account_id,omitempty" xml:"account_id,omitempty"`
}
