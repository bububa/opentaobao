package wdk

import (
	"sync"
)

// CouponQrcodeResultDo 结构体
type CouponQrcodeResultDo struct {
	// 二维码url
	QrcodeUrl string `json:"qrcode_url,omitempty" xml:"qrcode_url,omitempty"`
}

var poolCouponQrcodeResultDo = sync.Pool{
	New: func() any {
		return new(CouponQrcodeResultDo)
	},
}

// GetCouponQrcodeResultDo() 从对象池中获取CouponQrcodeResultDo
func GetCouponQrcodeResultDo() *CouponQrcodeResultDo {
	return poolCouponQrcodeResultDo.Get().(*CouponQrcodeResultDo)
}

// ReleaseCouponQrcodeResultDo 释放CouponQrcodeResultDo
func ReleaseCouponQrcodeResultDo(v *CouponQrcodeResultDo) {
	v.QrcodeUrl = ""
	poolCouponQrcodeResultDo.Put(v)
}
