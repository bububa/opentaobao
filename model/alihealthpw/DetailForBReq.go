package alihealthpw

import (
	"sync"
)

// DetailForBReq 结构体
type DetailForBReq struct {
	// 申请单唯一编码
	ApplyUniqueCode string `json:"apply_unique_code,omitempty" xml:"apply_unique_code,omitempty"`
}

var poolDetailForBReq = sync.Pool{
	New: func() any {
		return new(DetailForBReq)
	},
}

// GetDetailForBReq() 从对象池中获取DetailForBReq
func GetDetailForBReq() *DetailForBReq {
	return poolDetailForBReq.Get().(*DetailForBReq)
}

// ReleaseDetailForBReq 释放DetailForBReq
func ReleaseDetailForBReq(v *DetailForBReq) {
	v.ApplyUniqueCode = ""
	poolDetailForBReq.Put(v)
}
