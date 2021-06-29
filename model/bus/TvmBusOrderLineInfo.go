package bus

// TvmBusOrderLineInfo 
type TvmBusOrderLineInfo struct {
    // 代理商订单号
    AgentOrderId   string `json:"agent_order_id,omitempty" xml:"agent_order_id,omitempty"`
    // 阿里支付交易号
    AlipayTradeNo   string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
    // 阿里订单编号
    AlitripOrderId   string `json:"alitrip_order_id,omitempty" xml:"alitrip_order_id,omitempty"`
    // 订单创建时间
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // 出票时间
    IssueTime   string `json:"issue_time,omitempty" xml:"issue_time,omitempty"`
    // 订单状态：    CREATED_NO_PAY(10, "已创建，待支付"),      PAYED_NO_NOTIFY(20, "已支付，待通知代理商"),      PAYED_AND_NOTIFY(30, "已支付，待代理商出票"),//bookorder      BOOKED_AND_CONFIRM(50, "已出票，已确认"),//booksucess     BOOK_FAILED(-1, "出票失败"),     PAY_TIME_OUT(-2, "支付超时"),     BOOK_TIME_OUT(-3, "出票超时"),     CANCEL(-4, "订单取消"),     CLOSE(-5, "订单关闭");
    OrderStatus   int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
    // passengers
    Passengers   []TvmPassengerVo `json:"passengers,omitempty" xml:"passengers>tvm_passenger_vo,omitempty"`
    // 票数
    TicketCount   int64 `json:"ticket_count,omitempty" xml:"ticket_count,omitempty"`
    // totalPrice
    TotalPrice   int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
    // tvmBusLineInfo
    TvmBusLineInfo   *TvmBusLineInfo `json:"tvm_bus_line_info,omitempty" xml:"tvm_bus_line_info,omitempty"`
    // 支付时间
    PayTime   string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
    // 退款信息
    Refunds   []TvmRefundApply `json:"refunds,omitempty" xml:"refunds>tvm_refund_apply,omitempty"`
}
