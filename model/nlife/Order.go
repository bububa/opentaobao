package nlife

// Order 
type Order struct {
    // 商品列表
    GoodsList   []Goods `json:"goods_list,omitempty" xml:"goods_list>goods,omitempty"`
    // 收银员姓名
    SalesName   string `json:"sales_name,omitempty" xml:"sales_name,omitempty"`
    // 交易状态:WAIT_PAY：等待支付；SUCCESS：支付成功；CLOSED：交易关闭
    TradeStatus   string `json:"trade_status,omitempty" xml:"trade_status,omitempty"`
    // 支付时间
    GmtPay   string `json:"gmt_pay,omitempty" xml:"gmt_pay,omitempty"`
    // 实付金额，人民币 分
    ActualPayAmount   int64 `json:"actual_pay_amount,omitempty" xml:"actual_pay_amount,omitempty"`
    // 订单金额，人民币 分
    TotalAmount   int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
    // 下单时间
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // 门店名称
    StoreName   string `json:"store_name,omitempty" xml:"store_name,omitempty"`
    // 门店ID
    StoreId   int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
    // 零售+订单号
    TradeNo   string `json:"trade_no,omitempty" xml:"trade_no,omitempty"`
    // 支付渠道
    PayChannel   string `json:"pay_channel,omitempty" xml:"pay_channel,omitempty"`
    // 已退金额，人民币 分
    RefundedAmount   int64 `json:"refunded_amount,omitempty" xml:"refunded_amount,omitempty"`
    // 退款状态：REFUNDED_PART：已部分退款；REFUNDED：已全部退款
    RefundStatus   string `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
    // 发生退货的商品列表
    RefundedGoods   []Goods `json:"refunded_goods,omitempty" xml:"refunded_goods>goods,omitempty"`
    // 收银员ID
    SalesId   string `json:"sales_id,omitempty" xml:"sales_id,omitempty"`
    // 0：门店订单；1：全渠道订单；3：网直供订单
    TradeBizType   int64 `json:"trade_biz_type,omitempty" xml:"trade_biz_type,omitempty"`
    // 如果是全渠道订单，此处为淘宝订单号
    OmniTradeNo   string `json:"omni_trade_no,omitempty" xml:"omni_trade_no,omitempty"`
    // 外部订单号 即业务方订单号
    OutTradeNo   string `json:"out_trade_no,omitempty" xml:"out_trade_no,omitempty"`
}
