package alitripmerchant

// ActivityDrawGoodsPostInfoVo 结构体
type ActivityDrawGoodsPostInfoVo struct {
	// 收货人地址
	ReceiverAddress string `json:"receiver_address,omitempty" xml:"receiver_address,omitempty"`
	// 收货人电话
	ReceiverPhone string `json:"receiver_phone,omitempty" xml:"receiver_phone,omitempty"`
	// 收货人姓名
	ReceiverName string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
}
