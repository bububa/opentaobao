package wdk

// RefundChannelVo 
/* model for simplify = false
type RefundChannelVo struct {

    // 可退金额
    
    Amount   int64 `json:"amount,omitempty"`
    

    // 渠道码
    
    Code   string `json:"code,omitempty"`
    

    // 渠道名
    
    Name   string `json:"name,omitempty"`
    

    // 付款吗
    
    PayCode   string `json:"pay_code,omitempty"`
    

    // 已退金额 (单位分)
    
    RefundedAmount   int64 `json:"refunded_amount,omitempty"`
    

    // 渠道退款状态
    
    Status   int64 `json:"status,omitempty"`
    

    // 渠道幂等ID
    
    UniqueId   string `json:"unique_id,omitempty"`
    

}
*/

// RefundChannelVo 
type RefundChannelVo struct {

    // 可退金额
    Amount   int64 `json:"amount,omitempty"`

    // 渠道码
    Code   string `json:"code,omitempty"`

    // 渠道名
    Name   string `json:"name,omitempty"`

    // 付款吗
    PayCode   string `json:"pay_code,omitempty"`

    // 已退金额 (单位分)
    RefundedAmount   int64 `json:"refunded_amount,omitempty"`

    // 渠道退款状态
    Status   int64 `json:"status,omitempty"`

    // 渠道幂等ID
    UniqueId   string `json:"unique_id,omitempty"`

}
