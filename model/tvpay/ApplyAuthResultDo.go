package tvpay

// ApplyAuthResultDo 结构体
type ApplyAuthResultDo struct {
	// 授权方式
	AuthMode string `json:"auth_mode,omitempty" xml:"auth_mode,omitempty"`
	// 手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 二维码地址
	QrCodeUrl string `json:"qr_code_url,omitempty" xml:"qr_code_url,omitempty"`
	// qrcode
	QrCode string `json:"qr_code,omitempty" xml:"qr_code,omitempty"`
}
