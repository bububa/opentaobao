package tvpay

import (
	"sync"
)

// GetPartnerPayResultResultDo 结构体
type GetPartnerPayResultResultDo struct {
	// 加密串，订单详情
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}

var poolGetPartnerPayResultResultDo = sync.Pool{
	New: func() any {
		return new(GetPartnerPayResultResultDo)
	},
}

// GetGetPartnerPayResultResultDo() 从对象池中获取GetPartnerPayResultResultDo
func GetGetPartnerPayResultResultDo() *GetPartnerPayResultResultDo {
	return poolGetPartnerPayResultResultDo.Get().(*GetPartnerPayResultResultDo)
}

// ReleaseGetPartnerPayResultResultDo 释放GetPartnerPayResultResultDo
func ReleaseGetPartnerPayResultResultDo(v *GetPartnerPayResultResultDo) {
	v.Data = ""
	poolGetPartnerPayResultResultDo.Put(v)
}
