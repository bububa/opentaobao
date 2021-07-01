package qimen

// TaobaoQimenReceiverinfoQueryResponse 
type TaobaoQimenReceiverinfoQueryResponse struct {
    // success|failure，必填
    Flag   string `json:"flag,omitempty" xml:"flag,omitempty"`
    // 响应码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 响应信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 订单收件人 ID, string (50)
    Oaid   string `json:"oaid,omitempty" xml:"oaid,omitempty"`
    // 出库单号, string (50) , 必填
    DeliveryOrderCode   string `json:"deliveryOrderCode,omitempty" xml:"deliveryOrderCode,omitempty"`
    // 收货人信息
    ReceiverInfo   *ReceiverInfo `json:"receiverInfo,omitempty" xml:"receiverInfo,omitempty"`
}
