package util

import (
	"sync"
)

// AlibabaMosFalconPosCounterQueryResultDo 结构体
type AlibabaMosFalconPosCounterQueryResultDo struct {
	// 标题
	Titles []string `json:"titles,omitempty" xml:"titles>string,omitempty"`
	// traceId
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 额外
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// total
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 是否成功
	Data *PosInfoDto `json:"data,omitempty" xml:"data,omitempty"`
	// 错误码
	ErrCode int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaMosFalconPosCounterQueryResultDo = sync.Pool{
	New: func() any {
		return new(AlibabaMosFalconPosCounterQueryResultDo)
	},
}

// GetAlibabaMosFalconPosCounterQueryResultDo() 从对象池中获取AlibabaMosFalconPosCounterQueryResultDo
func GetAlibabaMosFalconPosCounterQueryResultDo() *AlibabaMosFalconPosCounterQueryResultDo {
	return poolAlibabaMosFalconPosCounterQueryResultDo.Get().(*AlibabaMosFalconPosCounterQueryResultDo)
}

// ReleaseAlibabaMosFalconPosCounterQueryResultDo 释放AlibabaMosFalconPosCounterQueryResultDo
func ReleaseAlibabaMosFalconPosCounterQueryResultDo(v *AlibabaMosFalconPosCounterQueryResultDo) {
	v.Titles = v.Titles[:0]
	v.TraceId = ""
	v.Extra = ""
	v.ErrMsg = ""
	v.Total = 0
	v.Data = nil
	v.ErrCode = 0
	v.Success = false
	poolAlibabaMosFalconPosCounterQueryResultDo.Put(v)
}
