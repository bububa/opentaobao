package wlb

// OrderItems 
type OrderItems struct {

    // 明细ID
    OrderItemId   int64 `json:"order_item_id,omitempty"`

    // 明细对应主单的交易单号
    TradeId   string `json:"trade_id,omitempty"`

    // 明细对应的子交易单号
    TradeItemId   string `json:"trade_item_id,omitempty"`

    // 默认：0；促销赠品1001
    ItemTag   string `json:"item_tag,omitempty"`

    // 货品id
    ScItemId   string `json:"sc_item_id,omitempty"`

    // 前端商家编码
    ItemCode   string `json:"item_code,omitempty"`

    // 数量
    ItemQuantity   int64 `json:"item_quantity,omitempty"`

    // 商品金额 123.33元，单位：分
    ItemAmount   int64 `json:"item_amount,omitempty"`

    // 前端宝贝itemId
    ItemId   string `json:"item_id,omitempty"`

    // 前端skuId
    SkuId   string `json:"sku_id,omitempty"`

    // 后端商家编码
    ScItemCode   string `json:"sc_item_code,omitempty"`

}
