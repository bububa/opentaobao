package alitripmerchant

// AlitripmerchantgalaxyderbymembervouchercardshowqrcodeResponse 结构体
type AlitripmerchantgalaxyderbymembervouchercardshowqrcodeResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 500
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 1com.fliggy.merchant.unicorn.api.derbyVoucherCard.vo.DerbyVoucherIdentityQrcodeVO
	Content *DerbyVoucherIdentityQrcodeVo `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
