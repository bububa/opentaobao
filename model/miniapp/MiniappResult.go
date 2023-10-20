package miniapp

import (
	"sync"
)

// MiniappResult 结构体
type MiniappResult struct {
	// model
	Model []AppChannelConfigDto `json:"model,omitempty" xml:"model>app_channel_config_dto,omitempty"`
	// 错误码
	ErrorType string `json:"error_type,omitempty" xml:"error_type,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrType int64 `json:"err_type,omitempty" xml:"err_type,omitempty"`
	// true or false
	Successful bool `json:"successful,omitempty" xml:"successful,omitempty"`
}

var poolMiniappResult = sync.Pool{
	New: func() any {
		return new(MiniappResult)
	},
}

// GetMiniappResult() 从对象池中获取MiniappResult
func GetMiniappResult() *MiniappResult {
	return poolMiniappResult.Get().(*MiniappResult)
}

// ReleaseMiniappResult 释放MiniappResult
func ReleaseMiniappResult(v *MiniappResult) {
	v.Model = v.Model[:0]
	v.ErrorType = ""
	v.ErrorMsg = ""
	v.ErrType = 0
	v.Successful = false
	poolMiniappResult.Put(v)
}
