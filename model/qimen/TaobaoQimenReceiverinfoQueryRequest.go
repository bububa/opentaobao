package qimen

// TaobaoQimenReceiverinfoQueryRequest 
type TaobaoQimenReceiverinfoQueryRequest struct {
    // 订单收件人 ID, string (50)
    Oaid   string `json:"oaid,omitempty" xml:"oaid,omitempty"`
    // 出库单号, string (50) , 必填
    DeliveryOrderCode   string `json:"deliveryOrderCode,omitempty" xml:"deliveryOrderCode,omitempty"`
    // 货主ID
    OwnerCode   string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
}
