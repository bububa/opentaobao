package trade

// FastBuyPosReverseRequest 
type FastBuyPosReverseRequest struct {
    // 外部唯一订单号
    OutOrderId   string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
    // 外部退款单号
    OutRefundId   string `json:"out_refund_id,omitempty" xml:"out_refund_id,omitempty"`
    // 经营店id
    StoreId   string `json:"store_id,omitempty" xml:"store_id,omitempty"`
    // 退款操作员工工号
    OperatorNum   string `json:"operator_num,omitempty" xml:"operator_num,omitempty"`
    // 逆向支付宝交易号，对账使用
    RefundAlipayTradeId   string `json:"refund_alipay_trade_id,omitempty" xml:"refund_alipay_trade_id,omitempty"`
    // 外部门店编码
    OutShopCode   string `json:"out_shop_code,omitempty" xml:"out_shop_code,omitempty"`
}
