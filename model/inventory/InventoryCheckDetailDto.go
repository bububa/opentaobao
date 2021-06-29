package inventory

// InventoryCheckDetailDTO 
type InventoryCheckDetailDTO struct {
    // 如果是门店类型,则为必填。 ONLINE_INVENTORY  线上可售库存，  SHARE_INVENTORY 线下可售库存
    InvBizCode   string `json:"inv_biz_code,omitempty" xml:"inv_biz_code,omitempty"`
    // 调整数量，正数盘盈，负数盘亏
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // 若为货品仓库存，则此处是货品ID 若为商品直接设置仓库存，则此处是商品ID， 若商品带SKU，还需要补充skuId
    ScItemId   int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
    // 每个货品的调整子单据号，作为业务调整依据，处理时会根据此单据号作幂
    SubOrderId   string `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
    // 调整商品对应的SKUID，如果商品为货品，则为0
    SkuId   int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}
