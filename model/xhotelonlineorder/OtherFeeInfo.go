package xhotelonlineorder

// OtherFeeInfo 结构体
type OtherFeeInfo struct {
	// 杂费名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 杂费金额
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 结账状态(1:结账,0:未结账)
	Checkout int64 `json:"checkout,omitempty" xml:"checkout,omitempty"`
}
