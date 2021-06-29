package idle

// RecycleOrderSynDto 
type RecycleOrderSynDto struct {
    // 属性入参
    Attribute   *Attribute `json:"attribute,omitempty" xml:"attribute,omitempty"`
    // 订单号
    BizOrderId   string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
    // 1:置订单可见 2:已上门取件    3:已质检      4:卖家确认交易完成      5:回收商确认交易完成      6:卖家订单已评价      7:回收商订单已评价  8:信用预付订单打款 101:货物已退回  103:回收商关闭订单  104:支付宝代扣失败 105:支付宝代扣成 106:支付宝代扣逾期
    OrderStatus   string `json:"order_status,omitempty" xml:"order_status,omitempty"`
    // 回收商在top上申请的appkey
    PartnerKey   string `json:"partner_key,omitempty" xml:"partner_key,omitempty"`
}
