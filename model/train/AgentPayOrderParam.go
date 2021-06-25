package train

// AgentPayOrderParam 
type AgentPayOrderParam struct {

    // 改签单号
    ApplyId   int64 `json:"apply_id,omitempty"`

    // 12306票号
    SequenceNo   string `json:"sequence_no,omitempty"`

    // 12306支付url
    PayUrl   string `json:"pay_url,omitempty"`

    // 主订单号
    MainOrderId   int64 `json:"main_order_id,omitempty"`

}
