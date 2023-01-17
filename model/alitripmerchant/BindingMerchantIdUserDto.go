package alitripmerchant

// BindingMerchantIdUserDto 结构体
type BindingMerchantIdUserDto struct {
	// 1
	MerchantId string `json:"merchant_id,omitempty" xml:"merchant_id,omitempty"`
	// 1
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 1
	Type string `json:"type,omitempty" xml:"type,omitempty"`
}
