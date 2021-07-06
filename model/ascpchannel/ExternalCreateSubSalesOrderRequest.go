package ascpchannel

// ExternalCreateSubSalesOrderRequest 结构体
type ExternalCreateSubSalesOrderRequest struct {
	// 外部前端sku id
	OutSkuId string `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	// 外部前端宝贝id
	OutItemId string `json:"out_item_id,omitempty" xml:"out_item_id,omitempty"`
	// 外部子订单号, 外部一定要设置,如果外部没有,则设置为outOrderNo+productId+productSkuId
	OutSubOrderNo string `json:"out_sub_order_no,omitempty" xml:"out_sub_order_no,omitempty"`
	// 购买数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 产品skuId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 产品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}
