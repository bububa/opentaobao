package mos

// PosOrderSaleItemDTO 
type PosOrderSaleItemDTO struct {
    // 折扣金额
    DiscountAmount   int64 `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
    // 扩展信息
    ExtendParams   string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
    // 商品ID
    GoodsId   string `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
    // 开票行号
    GoodsLineNo   int64 `json:"goods_line_no,omitempty" xml:"goods_line_no,omitempty"`
    // 商品名称
    GoodsName   string `json:"goods_name,omitempty" xml:"goods_name,omitempty"`
    // 商品类型：  N：老单品；Y：喵货；C：云单品；M：零售+      *      ?? 非单品是啥？
    ItemType   string `json:"item_type,omitempty" xml:"item_type,omitempty"`
    // 门店号
    MallNo   string `json:"mall_no,omitempty" xml:"mall_no,omitempty"`
    // 原始价格
    OriPrice   int64 `json:"ori_price,omitempty" xml:"ori_price,omitempty"`
    // 应付金额
    PayAmount   int64 `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
    // 商品数量
    Quantity   string `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // 开票单号
    SaleTicketNo   string `json:"sale_ticket_no,omitempty" xml:"sale_ticket_no,omitempty"`
    // 结算码
    SettleCode   string `json:"settle_code,omitempty" xml:"settle_code,omitempty"`
    // 专柜号
    ShopNo   string `json:"shop_no,omitempty" xml:"shop_no,omitempty"`
}
