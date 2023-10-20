package ott

import (
	"sync"
)

// YunosTvscreenLgeLauncherGetResult 结构体
type YunosTvscreenLgeLauncherGetResult struct {
	// Error message when success == false
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// Error code when success == false
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// Metadata Info
	Model *ChannelDto `json:"model,omitempty" xml:"model,omitempty"`
	// httpStatusCode
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
	// Is process succeed.
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolYunosTvscreenLgeLauncherGetResult = sync.Pool{
	New: func() any {
		return new(YunosTvscreenLgeLauncherGetResult)
	},
}

// GetYunosTvscreenLgeLauncherGetResult() 从对象池中获取YunosTvscreenLgeLauncherGetResult
func GetYunosTvscreenLgeLauncherGetResult() *YunosTvscreenLgeLauncherGetResult {
	return poolYunosTvscreenLgeLauncherGetResult.Get().(*YunosTvscreenLgeLauncherGetResult)
}

// ReleaseYunosTvscreenLgeLauncherGetResult 释放YunosTvscreenLgeLauncherGetResult
func ReleaseYunosTvscreenLgeLauncherGetResult(v *YunosTvscreenLgeLauncherGetResult) {
	v.MsgInfo = ""
	v.MsgCode = ""
	v.Model = nil
	v.HttpStatusCode = 0
	v.Success = false
	poolYunosTvscreenLgeLauncherGetResult.Put(v)
}
