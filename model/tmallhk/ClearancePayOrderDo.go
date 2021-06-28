package tmallhk

// ClearancePayOrderDo 
/* model for simplify = false
type ClearancePayOrderDo struct {

    // 支付宝买家ID
    
    AlipayBuyerId   string `json:"alipay_buyer_id,omitempty"`
    

    // 支付单号
    
    PayOrderId   int64 `json:"pay_order_id,omitempty"`
    

}
*/

// ClearancePayOrderDo 
type ClearancePayOrderDo struct {

    // 支付宝买家ID
    AlipayBuyerId   string `json:"alipay_buyer_id,omitempty"`

    // 支付单号
    PayOrderId   int64 `json:"pay_order_id,omitempty"`

}
