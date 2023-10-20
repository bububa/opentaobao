package tmallcar

// ConfirmOrderQrCode4isvDto 结构体
type ConfirmOrderQrCode4isvDto struct {
	// 附加参数
	Extension string `json:"extension,omitempty" xml:"extension,omitempty"`
	// 确认订单页二维码
	QrCodeImgUrl string `json:"qr_code_img_url,omitempty" xml:"qr_code_img_url,omitempty"`
}
