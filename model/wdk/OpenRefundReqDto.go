package wdk

// OpenRefundReqDto 
type OpenRefundReqDto struct {

    // 经营店ID
    StoreId   string `json:"store_id,omitempty"`

    // 淘系子单单号
    SubBizOrderId   string `json:"sub_biz_order_id,omitempty"`

    // 退款操作人
    Agent   string `json:"agent,omitempty"`

    // 退款备注
    Memo   string `json:"memo,omitempty"`

    // 退款渠道
    RefundChannelList   []ChannelRefundDto `json:"refund_channel_list,omitempty"`

}
