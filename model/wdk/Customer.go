package wdk

// Customer 结构体
type Customer struct {
	// 收货人地址
	BuyerAddress string `json:"buyer_address,omitempty" xml:"buyer_address,omitempty"`
	// 收货人联系电话
	BuyerPhone string `json:"buyer_phone,omitempty" xml:"buyer_phone,omitempty"`
	// 收货人姓名
	BuyerName string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
}
