package wdk

// PosSubOrderDo 
type PosSubOrderDo struct {

    // 库存单位，必填
    StockUnit   string `json:"stock_unit,omitempty"`

    // 库存单位购买数量，必填
    BuyAmountStock   string `json:"buy_amount_stock,omitempty"`

    // sku编码，必填
    SkuCode   string `json:"sku_code,omitempty"`

    // 外部子订单号，全局唯一，子单和主单不能重复，可以包含字母
    OutOrderId   string `json:"out_order_id,omitempty"`

    // 子单实付金额，单位分
    PayFee   int64 `json:"pay_fee,omitempty"`

    // 销售单位购买数量。对于标品，和库存单位库存单位购买数量一样
    BuyAmountSale   int64 `json:"buy_amount_sale,omitempty"`

    // 销售单位
    SaleUnit   string `json:"sale_unit,omitempty"`

    // 子单原价金额，单位分
    OriginFee   int64 `json:"origin_fee,omitempty"`

    // 子单优惠金额，单位分
    DiscountFee   int64 `json:"discount_fee,omitempty"`

    // 商品单价，单位分
    SkuPrice   int64 `json:"sku_price,omitempty"`

    // 商品名称
    SkuName   string `json:"sku_name,omitempty"`

}
