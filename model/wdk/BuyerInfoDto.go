package wdk

// BuyerInfoDto 结构体
type BuyerInfoDto struct {
	// 买家姓名
	BuyerName string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// 买家电话号码
	BuyerPhone string `json:"buyer_phone,omitempty" xml:"buyer_phone,omitempty"`
	// 收货地址
	BuyerAddress string `json:"buyer_address,omitempty" xml:"buyer_address,omitempty"`
}
