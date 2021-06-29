package ascpffo

// FulfillmentOrderStatusUpdateRequest 
type FulfillmentOrderStatusUpdateRequest struct {
    // 事件名
    Event   string `json:"event,omitempty" xml:"event,omitempty"`
    // 事件发生时间 unix时间戳
    EventTime   int64 `json:"event_time,omitempty" xml:"event_time,omitempty"`
    // 事件投递来源
    Source   string `json:"source,omitempty" xml:"source,omitempty"`
    // 事件发生原因
    Reason   string `json:"reason,omitempty" xml:"reason,omitempty"`
    // 外部业务单号
    OutBizId   string `json:"out_biz_id,omitempty" xml:"out_biz_id,omitempty"`
    // 交易单号
    TradeOrderId   string `json:"trade_order_id,omitempty" xml:"trade_order_id,omitempty"`
    // 交易子单号列表
    TradeOrderItemList   []string `json:"trade_order_item_list,omitempty" xml:"trade_order_item_list>string,omitempty"`
    // 运单号
    TrackingNumber   string `json:"tracking_number,omitempty" xml:"tracking_number,omitempty"`
}
