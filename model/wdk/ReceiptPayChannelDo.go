package wdk

// ReceiptPayChannelDo 
type ReceiptPayChannelDo struct {

    // 序号
    Index   int64 `json:"index,omitempty"`

    // 付款金额
    PayAmount   int64 `json:"pay_amount,omitempty"`

    // 付款类型， 现金、支票、银行卡、支付宝、微信
    PayType   string `json:"pay_type,omitempty"`

    // 款机号
    PosNo   string `json:"pos_no,omitempty"`

    // 流水号
    SerialNum   string `json:"serial_num,omitempty"`

    // 门店号
    StoreId   string `json:"store_id,omitempty"`

    // 外部支付单号
    ChannelOrderId   string `json:"channel_order_id,omitempty"`

    // 付款码
    PayCode   string `json:"pay_code,omitempty"`

}
