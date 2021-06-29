package tmallchannel

// TopChannelSubPurchaseOrderCreateParam 
type TopChannelSubPurchaseOrderCreateParam struct {
    // 采购数量
    BuyQuantity   int64 `json:"buy_quantity,omitempty" xml:"buy_quantity,omitempty"`
    // 采购货品SKU ID
    ProductSkuId   int64 `json:"product_sku_id,omitempty" xml:"product_sku_id,omitempty"`
    // 采购货品ID
    ProductId   int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}
