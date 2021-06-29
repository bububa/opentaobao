package wdk

// ReverseSkuInfo 
type ReverseSkuInfo struct {
    // 仓内报损数量（退仓数量-实际入库数量）
    WarehouseLossStockQuantity   string `json:"warehouse_loss_stock_quantity,omitempty" xml:"warehouse_loss_stock_quantity,omitempty"`
    // 实际入库数量(库存单位)
    ActualInBoundStockQuantity   string `json:"actual_in_bound_stock_quantity,omitempty" xml:"actual_in_bound_stock_quantity,omitempty"`
    // 商品编码
    SkuCode   string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
    // 逆向履约单号
    ReverseFulfillOrderId   string `json:"reverse_fulfill_order_id,omitempty" xml:"reverse_fulfill_order_id,omitempty"`
    // 逆向履约子单号
    ReverseFulfillSubOrderId   string `json:"reverse_fulfill_sub_order_id,omitempty" xml:"reverse_fulfill_sub_order_id,omitempty"`
    // 仓内报损原因列表
    WarehouseLossReasonList   []WarehouseLossReason `json:"warehouse_loss_reason_list,omitempty" xml:"warehouse_loss_reason_list>warehouse_loss_reason,omitempty"`
    // 关联的正向履约主单号
    RelatedFulfillOrderId   string `json:"related_fulfill_order_id,omitempty" xml:"related_fulfill_order_id,omitempty"`
    // 关联的正向履约子单号
    RelatedFulfillSubOrderId   string `json:"related_fulfill_sub_order_id,omitempty" xml:"related_fulfill_sub_order_id,omitempty"`
}
