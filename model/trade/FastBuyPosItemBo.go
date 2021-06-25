package trade

// FastBuyPosItemBo 
type FastBuyPosItemBo struct {

    // 商品条码
    BarCode   string `json:"bar_code,omitempty"`

    // 购买库存数量
    InvQuantity   string `json:"inv_quantity,omitempty"`

    // 商品行号,单次下单自增
    LineId   int64 `json:"line_id,omitempty"`

    // 一阶段优惠金额（商品实际优惠金额）：单价*数量-商家优惠
    PromotionFee   int64 `json:"promotion_fee,omitempty"`

    // 商家商品编码
    SkuCode   string `json:"sku_code,omitempty"`

    // 商品名称
    Title   string `json:"title,omitempty"`

    // 商品单价
    UnitPrice   int64 `json:"unit_price,omitempty"`

    // 扩展属性
    ExtendInfo   string `json:"extend_info,omitempty"`

}
