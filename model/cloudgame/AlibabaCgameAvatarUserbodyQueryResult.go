package cloudgame

import (
	"sync"
)

// AlibabaCgameAvatarUserbodyQueryResult 结构体
type AlibabaCgameAvatarUserbodyQueryResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 响应数据
	Data *AlibabaCgameAvatarUserbodyQueryData `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaCgameAvatarUserbodyQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaCgameAvatarUserbodyQueryResult)
	},
}

// GetAlibabaCgameAvatarUserbodyQueryResult() 从对象池中获取AlibabaCgameAvatarUserbodyQueryResult
func GetAlibabaCgameAvatarUserbodyQueryResult() *AlibabaCgameAvatarUserbodyQueryResult {
	return poolAlibabaCgameAvatarUserbodyQueryResult.Get().(*AlibabaCgameAvatarUserbodyQueryResult)
}

// ReleaseAlibabaCgameAvatarUserbodyQueryResult 释放AlibabaCgameAvatarUserbodyQueryResult
func ReleaseAlibabaCgameAvatarUserbodyQueryResult(v *AlibabaCgameAvatarUserbodyQueryResult) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	v.Success = false
	poolAlibabaCgameAvatarUserbodyQueryResult.Put(v)
}
