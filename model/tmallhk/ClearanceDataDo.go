package tmallhk

// ClearanceDataDO 
type ClearanceDataDO struct {
    // 订单数据封装
    BizOrderDO   *ClearanceBizOrderDO `json:"biz_order_d_o,omitempty" xml:"biz_order_d_o,omitempty"`
    // 支付单封装
    PayOrderDO   *ClearancePayOrderDO `json:"pay_order_d_o,omitempty" xml:"pay_order_d_o,omitempty"`
}
