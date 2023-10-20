package omniorder

// TaobaoomnidealerodersrefundaddressResult 结构体
type TaobaoomnidealerodersrefundaddressResult struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Success string `json:"success,omitempty" xml:"success,omitempty"`
	// 经销商订单退货地址
	Data *ScbRefundAddressDto `json:"data,omitempty" xml:"data,omitempty"`
}
