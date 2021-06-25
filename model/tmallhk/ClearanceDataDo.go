package tmallhk

// ClearanceDataDo 
type ClearanceDataDo struct {

    // 订单数据封装
    BizOrderDO   *ClearanceBizOrderDo `json:"biz_order_d_o,omitempty"`

    // 支付单封装
    PayOrderDO   *ClearancePayOrderDo `json:"pay_order_d_o,omitempty"`

}
