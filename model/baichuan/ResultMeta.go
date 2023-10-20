package baichuan

import (
	"sync"
)

// ResultMeta 结构体
type ResultMeta struct {
	// 返回码对应的文案
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 返回码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 返回的详细内容
	Data *ResultData `json:"data,omitempty" xml:"data,omitempty"`
}

var poolResultMeta = sync.Pool{
	New: func() any {
		return new(ResultMeta)
	},
}

// GetResultMeta() 从对象池中获取ResultMeta
func GetResultMeta() *ResultMeta {
	return poolResultMeta.Get().(*ResultMeta)
}

// ReleaseResultMeta 释放ResultMeta
func ReleaseResultMeta(v *ResultMeta) {
	v.Msg = ""
	v.Code = 0
	v.Data = nil
	poolResultMeta.Put(v)
}
