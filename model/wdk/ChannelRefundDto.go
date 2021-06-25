package wdk

// ChannelRefundDto 
type ChannelRefundDto struct {

    // 退款渠道编码
    ChannelCode   string `json:"channel_code,omitempty"`

    // 渠道对应的退款金额(单位分)
    RefundAmount   int64 `json:"refund_amount,omitempty"`

}
