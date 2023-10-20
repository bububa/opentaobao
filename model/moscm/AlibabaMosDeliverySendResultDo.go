package moscm

import (
	"sync"
)

// AlibabaMosDeliverySendResultDo 结构体
type AlibabaMosDeliverySendResultDo struct {
	// 异常信息
	SubMsg string `json:"sub_msg,omitempty" xml:"sub_msg,omitempty"`
	// 编码
	SubCode string `json:"sub_code,omitempty" xml:"sub_code,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaMosDeliverySendResultDo = sync.Pool{
	New: func() any {
		return new(AlibabaMosDeliverySendResultDo)
	},
}

// GetAlibabaMosDeliverySendResultDo() 从对象池中获取AlibabaMosDeliverySendResultDo
func GetAlibabaMosDeliverySendResultDo() *AlibabaMosDeliverySendResultDo {
	return poolAlibabaMosDeliverySendResultDo.Get().(*AlibabaMosDeliverySendResultDo)
}

// ReleaseAlibabaMosDeliverySendResultDo 释放AlibabaMosDeliverySendResultDo
func ReleaseAlibabaMosDeliverySendResultDo(v *AlibabaMosDeliverySendResultDo) {
	v.SubMsg = ""
	v.SubCode = ""
	v.Success = false
	poolAlibabaMosDeliverySendResultDo.Put(v)
}
