package alitripmerchant

// AlitripmerchantgalaxyderbymembervouchercardpurchasablequeryResponse 结构体
type AlitripmerchantgalaxyderbymembervouchercardpurchasablequeryResponse struct {
	// 1
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 1
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 可购权益卡实体类
	Content *DerbyVoucherCardPurchasableVo `json:"content,omitempty" xml:"content,omitempty"`
	// 1
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
