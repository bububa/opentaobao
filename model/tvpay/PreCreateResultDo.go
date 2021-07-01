package tvpay

// PreCreateResultDo 
type PreCreateResultDo struct {
    // 外部订单号
    OutOrderNo   string `json:"out_order_no,omitempty" xml:"out_order_no,omitempty"`
    // 二维码
    QrCode   string `json:"qr_code,omitempty" xml:"qr_code,omitempty"`
}
