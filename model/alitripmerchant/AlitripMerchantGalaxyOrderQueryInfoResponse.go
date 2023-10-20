package alitripmerchant

// AlitripMerchantGalaxyOrderQueryInfoResponse 结构体
type AlitripMerchantGalaxyOrderQueryInfoResponse struct {
	// 错误代码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息展示
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 订单详情页对象
	Content *VoucherOrderDetailVo `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
