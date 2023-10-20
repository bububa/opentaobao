package alicom

import (
	"sync"
)

// CommonRtnDo 结构体
type CommonRtnDo struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolCommonRtnDo = sync.Pool{
	New: func() any {
		return new(CommonRtnDo)
	},
}

// GetCommonRtnDo() 从对象池中获取CommonRtnDo
func GetCommonRtnDo() *CommonRtnDo {
	return poolCommonRtnDo.Get().(*CommonRtnDo)
}

// ReleaseCommonRtnDo 释放CommonRtnDo
func ReleaseCommonRtnDo(v *CommonRtnDo) {
	v.Message = ""
	v.Code = ""
	v.Data = false
	poolCommonRtnDo.Put(v)
}
