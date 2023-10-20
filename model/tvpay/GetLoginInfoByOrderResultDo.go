package tvpay

import (
	"sync"
)

// GetLoginInfoByOrderResultDo 结构体
type GetLoginInfoByOrderResultDo struct {
	// 登陆信息，json
	AccessData string `json:"access_data,omitempty" xml:"access_data,omitempty"`
	// 是否有登陆信息
	HasLoginInfo bool `json:"has_login_info,omitempty" xml:"has_login_info,omitempty"`
}

var poolGetLoginInfoByOrderResultDo = sync.Pool{
	New: func() any {
		return new(GetLoginInfoByOrderResultDo)
	},
}

// GetGetLoginInfoByOrderResultDo() 从对象池中获取GetLoginInfoByOrderResultDo
func GetGetLoginInfoByOrderResultDo() *GetLoginInfoByOrderResultDo {
	return poolGetLoginInfoByOrderResultDo.Get().(*GetLoginInfoByOrderResultDo)
}

// ReleaseGetLoginInfoByOrderResultDo 释放GetLoginInfoByOrderResultDo
func ReleaseGetLoginInfoByOrderResultDo(v *GetLoginInfoByOrderResultDo) {
	v.AccessData = ""
	v.HasLoginInfo = false
	poolGetLoginInfoByOrderResultDo.Put(v)
}
