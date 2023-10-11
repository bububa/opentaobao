package wdk

// WorkCallbackSubOrderInfo 结构体
type WorkCallbackSubOrderInfo struct {
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 业务子订单编码
	BizSubOrderId int64 `json:"biz_sub_order_id,omitempty" xml:"biz_sub_order_id,omitempty"`
	// 销售单位拣货数量 PICKED/PACKAGED状态必填
	PickSaleQuantity *BigDecimal `json:"pick_sale_quantity,omitempty" xml:"pick_sale_quantity,omitempty"`
	// 库存单位拣货数量 PICKED/PACKAGED状态必填
	PickStockQuantity *BigDecimal `json:"pick_stock_quantity,omitempty" xml:"pick_stock_quantity,omitempty"`
}
