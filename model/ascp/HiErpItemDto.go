package ascp

// HiErpItemDto 结构体
type HiErpItemDto struct {
	// 外部订单明细唯一标识
	OutOrderItemId string `json:"out_order_item_id,omitempty" xml:"out_order_item_id,omitempty"`
	// 后端货品编码， 货品id和货品编码二选一必填，优先取货品ID
	ScItemCode string `json:"sc_item_code,omitempty" xml:"sc_item_code,omitempty"`
	// 扩展字段，json格式
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 子交易单号
	SubTradeId string `json:"sub_trade_id,omitempty" xml:"sub_trade_id,omitempty"`
	// 后端货品ID， 货品id和货品编码二选一必填，优先取货品ID
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 货品金额
	ItemAmount int64 `json:"item_amount,omitempty" xml:"item_amount,omitempty"`
	// 货品数量
	ItemQty int64 `json:"item_qty,omitempty" xml:"item_qty,omitempty"`
}
