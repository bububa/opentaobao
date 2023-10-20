package alitripmerchant

// AlitripmerchantgalaxyderbymembervoucherofflineqrcodeResponse 结构体
type AlitripmerchantgalaxyderbymembervoucherofflineqrcodeResponse struct {
	// 200/500
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 权益券实体类
	Content *DerbyVoucherVo `json:"content,omitempty" xml:"content,omitempty"`
	// 成功/失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
