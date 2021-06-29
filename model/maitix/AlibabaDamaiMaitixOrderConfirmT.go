package maitix

// AlibabaDamaiMaitixOrderConfirmT 
type AlibabaDamaiMaitixOrderConfirmT struct {
    // 订单id
    OrderId   string `json:"order_id,omitempty" xml:"order_id,omitempty"`
    // 支付状态：0:失败,1:成功
    PayStatus   int64 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
}
