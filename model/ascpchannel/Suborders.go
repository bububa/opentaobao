package ascpchannel

// Suborders 结构体
type Suborders struct {
	// 外部skuId
	OutSkuId string `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	// 外部itemId
	OutItemId string `json:"out_item_id,omitempty" xml:"out_item_id,omitempty"`
	// 外部子订单号
	OutSubOrderNo string `json:"out_sub_order_no,omitempty" xml:"out_sub_order_no,omitempty"`
	// 外部订单号
	SubSaleOrderNo string `json:"sub_sale_order_no,omitempty" xml:"sub_sale_order_no,omitempty"`
	// skuId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 产品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}
