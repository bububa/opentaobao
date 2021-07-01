package lstlogistics

// LstThirdPartDetailShipOrderDto 
type LstThirdPartDetailShipOrderDto struct {
    // 销售数量
    SaleQuantity   int64 `json:"sale_quantity,omitempty" xml:"sale_quantity,omitempty"`
    // 销售单位
    SkuUnit   string `json:"sku_unit,omitempty" xml:"sku_unit,omitempty"`
    // 规格
    SkuSpec   string `json:"sku_spec,omitempty" xml:"sku_spec,omitempty"`
    // 条码
    SkuCode   string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
    // 商品名称
    SkuName   string `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
    // 装车数量
    LoadQuantity   int64 `json:"load_quantity,omitempty" xml:"load_quantity,omitempty"`
    // 签收数量
    SignQuantity   int64 `json:"sign_quantity,omitempty" xml:"sign_quantity,omitempty"`
    // 明细订单ID
    DetailOrderId   string `json:"detail_order_id,omitempty" xml:"detail_order_id,omitempty"`
    // 子发货单状态
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
}
