package wdk

// PosRefundCreateRequest 
type PosRefundCreateRequest struct {
    // 外部主订单号
    OutOrderId   string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
    // 渠道店id
    ShopId   string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
    // 经营店code
    StoreId   string `json:"store_id,omitempty" xml:"store_id,omitempty"`
    // 外部子订单号
    OutSubOrderId   string `json:"out_sub_order_id,omitempty" xml:"out_sub_order_id,omitempty"`
    // skuCode
    SkuCode   string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
    // 库存单位退货数量
    RefundAmountSale   int64 `json:"refund_amount_sale,omitempty" xml:"refund_amount_sale,omitempty"`
    // 销售单位退货数量
    RefundAmountStock   string `json:"refund_amount_stock,omitempty" xml:"refund_amount_stock,omitempty"`
    // 库存单位
    StockUnit   string `json:"stock_unit,omitempty" xml:"stock_unit,omitempty"`
    // 销售单位
    SaleUnit   string `json:"sale_unit,omitempty" xml:"sale_unit,omitempty"`
    // 退款金额，单位分
    RefundFee   int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
    // 退款类型 1：仅退款 2：退货退款
    RefundType   int64 `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
    // 退款速度类型 1：闪退,未离开货架  2：标准，离开货架
    RefundSpeedType   int64 `json:"refund_speed_type,omitempty" xml:"refund_speed_type,omitempty"`
    // 是否称重商品
    WeightItem   bool `json:"weight_item,omitempty" xml:"weight_item,omitempty"`
}
