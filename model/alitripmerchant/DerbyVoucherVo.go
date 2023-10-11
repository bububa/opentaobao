package alitripmerchant

// DerbyVoucherVo 结构体
type DerbyVoucherVo struct {
	// 线下权益券二维码
	QRCodeImage string `json:"q_r_code_image,omitempty" xml:"q_r_code_image,omitempty"`
	// 线下核销是否成功
	OfflineRedemption bool `json:"offline_redemption,omitempty" xml:"offline_redemption,omitempty"`
}
