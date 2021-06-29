package tmallchannel

// TopChannelApplyOrderRelateItemModifyParam 
type TopChannelApplyOrderRelateItemModifyParam struct {
    // 采购价
    PurchasePrice   int64 `json:"purchase_price,omitempty" xml:"purchase_price,omitempty"`
    // 购买数量
    BuyQuantity   int64 `json:"buy_quantity,omitempty" xml:"buy_quantity,omitempty"`
    // 对应货品的skuId
    ProductSkuId   int64 `json:"product_sku_id,omitempty" xml:"product_sku_id,omitempty"`
    // 对应的货品ID
    ProductId   int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}
