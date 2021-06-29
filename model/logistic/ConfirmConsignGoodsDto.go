package logistic

// ConfirmConsignGoodsDTO 
type ConfirmConsignGoodsDTO struct {
    // 待发货商品的前端宝贝id
    ItemId   string `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 待发货商品的数量
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // 待发货商品的子交易号
    TcSubTradeId   int64 `json:"tc_sub_trade_id,omitempty" xml:"tc_sub_trade_id,omitempty"`
}
