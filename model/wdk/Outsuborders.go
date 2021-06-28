package wdk

// Outsuborders 
/* model for simplify = false
type Outsuborders struct {

    // 渠道子订单ID
    
    OutSubOrderId   string `json:"out_sub_order_id,omitempty"`
    

    // 渠道子订单最大可退金额
    
    MaxRefundFee   int64 `json:"max_refund_fee,omitempty"`
    

    // 当逆向已达到整单退时, 最后一个会有退运费金额
    
    PostFee   int64 `json:"post_fee,omitempty"`
    

    // 是否可退
    
    CanReverse   bool `json:"can_reverse,omitempty"`
    

}
*/

// Outsuborders 
type Outsuborders struct {

    // 渠道子订单ID
    OutSubOrderId   string `json:"out_sub_order_id,omitempty"`

    // 渠道子订单最大可退金额
    MaxRefundFee   int64 `json:"max_refund_fee,omitempty"`

    // 当逆向已达到整单退时, 最后一个会有退运费金额
    PostFee   int64 `json:"post_fee,omitempty"`

    // 是否可退
    CanReverse   bool `json:"can_reverse,omitempty"`

}
