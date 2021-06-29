package omniorder

// TradeOrderInfoDTO 
type TradeOrderInfoDTO struct {
    // 主订单ID
    MainOrderId   string `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
    // 子订单信息
    SubOrders   []TradeOrderDetailDTO `json:"sub_orders,omitempty" xml:"sub_orders>trade_order_detail_dto,omitempty"`
}
