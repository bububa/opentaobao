package jipiao

// TripConfirmOrderParam 
type TripConfirmOrderParam struct {
    // 确认项：官网投放卖家确认订单参数为orderConfirmHk
    Confirm   string `json:"confirm,omitempty" xml:"confirm,omitempty"`
    // 对于此订单的描述。
    Info   string `json:"info,omitempty" xml:"info,omitempty"`
    // 订单号
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}
