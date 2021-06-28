package tmallhk

// ClearanceDataDo 
/* model for simplify = false
type ClearanceDataDo struct {

    // 订单数据封装
    
    BizOrderDO  *struct {
        ClearanceBizOrderDo  *ClearanceBizOrderDo `json:"clearance_biz_order_do,omitempty"`
    } `json:"biz_order_d_o,omitempty"`
    

    // 支付单封装
    
    PayOrderDO  *struct {
        ClearancePayOrderDo  *ClearancePayOrderDo `json:"clearance_pay_order_do,omitempty"`
    } `json:"pay_order_d_o,omitempty"`
    

}
*/

// ClearanceDataDo 
type ClearanceDataDo struct {

    // 订单数据封装
    BizOrderDO   *ClearanceBizOrderDo `json:"biz_order_d_o,omitempty"`

    // 支付单封装
    PayOrderDO   *ClearancePayOrderDo `json:"pay_order_d_o,omitempty"`

}
