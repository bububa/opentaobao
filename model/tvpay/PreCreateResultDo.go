package tvpay

import (
	"sync"
)

// PreCreateResultDo 结构体
type PreCreateResultDo struct {
	// 外部订单号
	OutOrderNo string `json:"out_order_no,omitempty" xml:"out_order_no,omitempty"`
	// 二维码
	QrCode string `json:"qr_code,omitempty" xml:"qr_code,omitempty"`
}

var poolPreCreateResultDo = sync.Pool{
	New: func() any {
		return new(PreCreateResultDo)
	},
}

// GetPreCreateResultDo() 从对象池中获取PreCreateResultDo
func GetPreCreateResultDo() *PreCreateResultDo {
	return poolPreCreateResultDo.Get().(*PreCreateResultDo)
}

// ReleasePreCreateResultDo 释放PreCreateResultDo
func ReleasePreCreateResultDo(v *PreCreateResultDo) {
	v.OutOrderNo = ""
	v.QrCode = ""
	poolPreCreateResultDo.Put(v)
}
