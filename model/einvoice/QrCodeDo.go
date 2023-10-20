package einvoice

import (
	"sync"
)

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

var poolQrCodeDo = sync.Pool{
	New: func() any {
		return new(QrCodeDo)
	},
}

// GetQrCodeDo() 从对象池中获取QrCodeDo
func GetQrCodeDo() *QrCodeDo {
	return poolQrCodeDo.Get().(*QrCodeDo)
}

// ReleaseQrCodeDo 释放QrCodeDo
func ReleaseQrCodeDo(v *QrCodeDo) {
	v.QrLogo = ""
	v.QrType = 0
	v.Width = 0
	v.Height = 0
	poolQrCodeDo.Put(v)
}
