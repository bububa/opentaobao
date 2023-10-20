package alimember

import (
	"sync"
)

// AlibabaMemberPointOperateTopResult 结构体
type AlibabaMemberPointOperateTopResult struct {
	// 返回码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// json格式
	Data *PointOperateResponseDto `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaMemberPointOperateTopResult = sync.Pool{
	New: func() any {
		return new(AlibabaMemberPointOperateTopResult)
	},
}

// GetAlibabaMemberPointOperateTopResult() 从对象池中获取AlibabaMemberPointOperateTopResult
func GetAlibabaMemberPointOperateTopResult() *AlibabaMemberPointOperateTopResult {
	return poolAlibabaMemberPointOperateTopResult.Get().(*AlibabaMemberPointOperateTopResult)
}

// ReleaseAlibabaMemberPointOperateTopResult 释放AlibabaMemberPointOperateTopResult
func ReleaseAlibabaMemberPointOperateTopResult(v *AlibabaMemberPointOperateTopResult) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	poolAlibabaMemberPointOperateTopResult.Put(v)
}
