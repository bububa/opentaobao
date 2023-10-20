package cloudgame

import (
	"sync"
)

// AlibabaCloudgameInteractiveGameStartResult 结构体
type AlibabaCloudgameInteractiveGameStartResult struct {
	// 返回状态码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Data *StartGameResponse `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaCloudgameInteractiveGameStartResult = sync.Pool{
	New: func() any {
		return new(AlibabaCloudgameInteractiveGameStartResult)
	},
}

// GetAlibabaCloudgameInteractiveGameStartResult() 从对象池中获取AlibabaCloudgameInteractiveGameStartResult
func GetAlibabaCloudgameInteractiveGameStartResult() *AlibabaCloudgameInteractiveGameStartResult {
	return poolAlibabaCloudgameInteractiveGameStartResult.Get().(*AlibabaCloudgameInteractiveGameStartResult)
}

// ReleaseAlibabaCloudgameInteractiveGameStartResult 释放AlibabaCloudgameInteractiveGameStartResult
func ReleaseAlibabaCloudgameInteractiveGameStartResult(v *AlibabaCloudgameInteractiveGameStartResult) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	poolAlibabaCloudgameInteractiveGameStartResult.Put(v)
}
