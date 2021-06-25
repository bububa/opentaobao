package drug

// SubOrderDto 
type SubOrderDto struct {

    // 子订单ID
    SubOrderId   int64 `json:"sub_order_id,omitempty"`

    // 商品ID
    ItemId   int64 `json:"item_id,omitempty"`

    // 商品营销类型：0-普通，1-满就送；2-兑换券
    Type   int64 `json:"type,omitempty"`

    // 价格
    Price   int64 `json:"price,omitempty"`

    // 购买数量
    BuyAmount   int64 `json:"buy_amount,omitempty"`

    // 商品单位
    Unit   string `json:"unit,omitempty"`

    // 商品名称
    Title   string `json:"title,omitempty"`

    // 商家自定义商品ID
    OutItemdId   string `json:"out_itemd_id,omitempty"`

    // 是否为处方药
    Rx   int64 `json:"rx,omitempty"`

    // 套装商品子商品系信息
    SuitSubItemDtoList   []SuitSubItemDto `json:"suit_sub_item_dto_list,omitempty"`

}
