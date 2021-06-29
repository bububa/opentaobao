package idle

// ConsignmentOrderSynDTO 
type ConsignmentOrderSynDTO struct {
    // 订单二级状态,一级状态的子状态,对于没有二级状态的场景该字段为空。一级状态为2已取件: 21:已取件; 22:已收件; 一级状态为3已质检: 31:已质检; 32:用户已确认; 201:一次挂拍; 一级状态为20竞拍中: 202:一次竞拍中; 203:一次竞拍成交; 204:一次拍卖违约; 205:一次竞拍流拍; 211:二次挂拍; 212:二次竞拍中; 213:二次竞拍成交; 214:二次拍卖违约; 215:二次竞拍流拍; 一级状态为5服务商确认交易完成: 51:拍卖成功/订单成功; 58:回收商确认交易/拍卖流拍成交; 59:服务商(兜底)确认交易/支付;
    OrderSubStatus   string `json:"order_sub_status,omitempty" xml:"order_sub_status,omitempty"`
    // 订单一级状态。1:已下单; 2:已取件; 3:已质检; 20:竞拍中; 5:交易成功; 6:卖家已评价; 7:服务商已评价; 100:申请退货; 101:已退货; 102:卖家取消订单关闭; 103:服务商取消订单关闭;
    OrderStatus   string `json:"order_status,omitempty" xml:"order_status,omitempty"`
    // 闲鱼订单号
    BizOrderId   string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
    // 履约节点数据
    Attribute   *Attribute `json:"attribute,omitempty" xml:"attribute,omitempty"`
}
