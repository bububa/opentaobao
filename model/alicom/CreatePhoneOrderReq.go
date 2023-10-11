package alicom

// CreatePhoneOrderReq 结构体
type CreatePhoneOrderReq struct {
	// 业务办理账户
	Account string `json:"account,omitempty" xml:"account,omitempty"`
	// 商家幂等ID
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}
