package cloudgame

import (
	"sync"
)

// AlibabaCgameMpMpsessionSendmessagetogameResult 结构体
type AlibabaCgameMpMpsessionSendmessagetogameResult struct {
	// 0
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// sucess
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// true
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaCgameMpMpsessionSendmessagetogameResult = sync.Pool{
	New: func() any {
		return new(AlibabaCgameMpMpsessionSendmessagetogameResult)
	},
}

// GetAlibabaCgameMpMpsessionSendmessagetogameResult() 从对象池中获取AlibabaCgameMpMpsessionSendmessagetogameResult
func GetAlibabaCgameMpMpsessionSendmessagetogameResult() *AlibabaCgameMpMpsessionSendmessagetogameResult {
	return poolAlibabaCgameMpMpsessionSendmessagetogameResult.Get().(*AlibabaCgameMpMpsessionSendmessagetogameResult)
}

// ReleaseAlibabaCgameMpMpsessionSendmessagetogameResult 释放AlibabaCgameMpMpsessionSendmessagetogameResult
func ReleaseAlibabaCgameMpMpsessionSendmessagetogameResult(v *AlibabaCgameMpMpsessionSendmessagetogameResult) {
	v.Code = ""
	v.Message = ""
	v.Data = false
	poolAlibabaCgameMpMpsessionSendmessagetogameResult.Put(v)
}
