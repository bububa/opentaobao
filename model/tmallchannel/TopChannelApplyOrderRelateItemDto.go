package tmallchannel

// TopChannelApplyOrderRelateItemDto 
type TopChannelApplyOrderRelateItemDto struct {
    // 货品ID
    ProductId   int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
    // 货品的skuId
    ProductSkuId   int64 `json:"product_sku_id,omitempty" xml:"product_sku_id,omitempty"`
    // 购买数量
    BuyQuantity   int64 `json:"buy_quantity,omitempty" xml:"buy_quantity,omitempty"`
    // 采购价
    PurchasePrice   int64 `json:"purchase_price,omitempty" xml:"purchase_price,omitempty"`
    // b2c的实付款
    B2cActualPayFee   int64 `json:"b2c_actual_pay_fee,omitempty" xml:"b2c_actual_pay_fee,omitempty"`
}
