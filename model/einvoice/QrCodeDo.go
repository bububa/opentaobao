package einvoice

// QrCodeDo 结构体
type QrCodeDo struct {
	// 二维码logo
	QrLogo string `json:"qr_logo,omitempty" xml:"qr_logo,omitempty"`
	// 二维码返回数据类型：1=二维码背后的URL，2=二维码图片CDN URL，3=二维码二进制数据流
	QrType int64 `json:"qr_type,omitempty" xml:"qr_type,omitempty"`
	// 二维码宽度
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
	// 二维码高度
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
}
