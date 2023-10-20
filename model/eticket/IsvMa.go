package eticket

import (
	"sync"
)

// IsvMa 结构体
type IsvMa struct {
	// 串码码值
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 二维码图片文件名。已经申请了上传二维码的码商必填，其它码商无需关心。这个值是taobao.eticket.merchant.img.upload调用后的file_name
	QrImage string `json:"qr_image,omitempty" xml:"qr_image,omitempty"`
	// 码的可核销份数
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
}

var poolIsvMa = sync.Pool{
	New: func() any {
		return new(IsvMa)
	},
}

// GetIsvMa() 从对象池中获取IsvMa
func GetIsvMa() *IsvMa {
	return poolIsvMa.Get().(*IsvMa)
}

// ReleaseIsvMa 释放IsvMa
func ReleaseIsvMa(v *IsvMa) {
	v.Code = ""
	v.QrImage = ""
	v.Num = 0
	poolIsvMa.Put(v)
}
