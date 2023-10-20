package cloudgame

import (
	"sync"
)

// AlibabaCgameLiteplayAvatarBodyQueryResult 结构体
type AlibabaCgameLiteplayAvatarBodyQueryResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回数据体
	Data *TopAvatarBodyDto `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaCgameLiteplayAvatarBodyQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaCgameLiteplayAvatarBodyQueryResult)
	},
}

// GetAlibabaCgameLiteplayAvatarBodyQueryResult() 从对象池中获取AlibabaCgameLiteplayAvatarBodyQueryResult
func GetAlibabaCgameLiteplayAvatarBodyQueryResult() *AlibabaCgameLiteplayAvatarBodyQueryResult {
	return poolAlibabaCgameLiteplayAvatarBodyQueryResult.Get().(*AlibabaCgameLiteplayAvatarBodyQueryResult)
}

// ReleaseAlibabaCgameLiteplayAvatarBodyQueryResult 释放AlibabaCgameLiteplayAvatarBodyQueryResult
func ReleaseAlibabaCgameLiteplayAvatarBodyQueryResult(v *AlibabaCgameLiteplayAvatarBodyQueryResult) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	poolAlibabaCgameLiteplayAvatarBodyQueryResult.Put(v)
}
