package ascpchannel

// ExternalCreateRefundOrderRequest 
type ExternalCreateRefundOrderRequest struct {
    // 币种
    CurrencyType   string `json:"currency_type,omitempty" xml:"currency_type,omitempty"`
    // 销售订单号
    SaleOrderNo   string `json:"sale_order_no,omitempty" xml:"sale_order_no,omitempty"`
    // 外部退款单号
    OutRefundNo   string `json:"out_refund_no,omitempty" xml:"out_refund_no,omitempty"`
    // 备注
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    // 外部订单号
    OutOrderNo   string `json:"out_order_no,omitempty" xml:"out_order_no,omitempty"`
}
