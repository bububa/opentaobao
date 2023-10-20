package tmallcar

import (
	"sync"
)

// ConfirmOrderQrCode4IsvDto 结构体
type ConfirmOrderQrCode4IsvDto struct {
	// 附加参数
	Extension string `json:"extension,omitempty" xml:"extension,omitempty"`
	// 确认订单页二维码
	QrCodeImgUrl string `json:"qr_code_img_url,omitempty" xml:"qr_code_img_url,omitempty"`
}

var poolConfirmOrderQrCode4IsvDto = sync.Pool{
	New: func() any {
		return new(ConfirmOrderQrCode4IsvDto)
	},
}

// GetConfirmOrderQrCode4IsvDto() 从对象池中获取ConfirmOrderQrCode4IsvDto
func GetConfirmOrderQrCode4IsvDto() *ConfirmOrderQrCode4IsvDto {
	return poolConfirmOrderQrCode4IsvDto.Get().(*ConfirmOrderQrCode4IsvDto)
}

// ReleaseConfirmOrderQrCode4IsvDto 释放ConfirmOrderQrCode4IsvDto
func ReleaseConfirmOrderQrCode4IsvDto(v *ConfirmOrderQrCode4IsvDto) {
	v.Extension = ""
	v.QrCodeImgUrl = ""
	poolConfirmOrderQrCode4IsvDto.Put(v)
}
