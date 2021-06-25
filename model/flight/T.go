package flight

// T 
type T struct {

    // 申请单号
    ApplyId   string `json:"apply_id,omitempty"`

    // 国内国际标识
    DomesticIntl   int64 `json:"domestic_intl,omitempty"`

    // 飞猪订单号
    OrderId   string `json:"order_id,omitempty"`

    // 申请时间
    ApplyTime   string `json:"apply_time,omitempty"`

    // 支付时间
    PayTime   string `json:"pay_time,omitempty"`

}
