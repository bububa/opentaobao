package bus

// TopInsMerchantInfo 结构体
type TopInsMerchantInfo struct {
	// 商户ID
	MerchantId string `json:"merchant_id,omitempty" xml:"merchant_id,omitempty"`
	// 商户名称
	MerchantName string `json:"merchant_name,omitempty" xml:"merchant_name,omitempty"`
}
