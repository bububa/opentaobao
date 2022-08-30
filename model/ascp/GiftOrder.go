package ascp

// GiftOrder 结构体
type GiftOrder struct {
	// 赠品绑赠单号，回滚时需要传递
	GiftOrderId string `json:"gift_order_id,omitempty" xml:"gift_order_id,omitempty"`
	// 货品商家编码
	ScItemCode string `json:"sc_item_code,omitempty" xml:"sc_item_code,omitempty"`
	// 货品名称
	ScItemName string `json:"sc_item_name,omitempty" xml:"sc_item_name,omitempty"`
	// 商品ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 活动ID
	ActivityId string `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 赠品ID
	GiftId string `json:"gift_id,omitempty" xml:"gift_id,omitempty"`
	// 卖家昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 活动名称
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 赠品数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}
