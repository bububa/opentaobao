package iot

// TradeItemDo 
type TradeItemDo struct {

    // 购买件数，默认为1
    BuyAmount   int64 `json:"buy_amount,omitempty"`

    // 商品SKU ID
    SkuId   int64 `json:"sku_id,omitempty"`

    // 商品ID
    ItemId   int64 `json:"item_id,omitempty"`

    // 售卖单价，单位分
    UnitPrice   int64 `json:"unit_price,omitempty"`

}
