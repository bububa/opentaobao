package cntms

// CnTmsLogisticsOrderItem 
type CnTmsLogisticsOrderItem struct {

    // ERP订单明细编码
    OrderItemId   string `json:"order_item_id,omitempty"`

    // 子交易号
    SubTradeId   string `json:"sub_trade_id,omitempty"`

    // 发货商品名称
    ItemName   string `json:"item_name,omitempty"`

    // 发货商品商品数量
    Quantity   int64 `json:"quantity,omitempty"`

    // 扩展字段 K:V;
    ExtendFields   string `json:"extend_fields,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

    // 商品单价，单位分
    ItemPrice   int64 `json:"item_price,omitempty"`

}
