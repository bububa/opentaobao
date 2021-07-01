package retail

// ItemLineDto 结构体
type ItemLineDto struct {
	// 百安居id
	OutLineId string `json:"out_line_id,omitempty" xml:"out_line_id,omitempty"`
	// 门店商品行id
	StoreOrderLineId int64 `json:"store_order_line_id,omitempty" xml:"store_order_line_id,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// skuId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 购买数量
	BuyAmount int64 `json:"buy_amount,omitempty" xml:"buy_amount,omitempty"`
	// 实际付款价格
	PayFee int64 `json:"pay_fee,omitempty" xml:"pay_fee,omitempty"`
	// 外部skuId
	OutSkuId string `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	// 服务子订单
	ServiceList []ServiceItemLineDto `json:"service_list,omitempty" xml:"service_list>service_item_line_dto,omitempty"`
}
