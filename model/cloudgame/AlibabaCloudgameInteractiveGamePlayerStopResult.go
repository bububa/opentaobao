package cloudgame

import (
	"sync"
)

// AlibabaCloudgameInteractiveGamePlayerStopResult 结构体
type AlibabaCloudgameInteractiveGamePlayerStopResult struct {
	// 返回状态码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Data *StopUserGameResponse `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaCloudgameInteractiveGamePlayerStopResult = sync.Pool{
	New: func() any {
		return new(AlibabaCloudgameInteractiveGamePlayerStopResult)
	},
}

// GetAlibabaCloudgameInteractiveGamePlayerStopResult() 从对象池中获取AlibabaCloudgameInteractiveGamePlayerStopResult
func GetAlibabaCloudgameInteractiveGamePlayerStopResult() *AlibabaCloudgameInteractiveGamePlayerStopResult {
	return poolAlibabaCloudgameInteractiveGamePlayerStopResult.Get().(*AlibabaCloudgameInteractiveGamePlayerStopResult)
}

// ReleaseAlibabaCloudgameInteractiveGamePlayerStopResult 释放AlibabaCloudgameInteractiveGamePlayerStopResult
func ReleaseAlibabaCloudgameInteractiveGamePlayerStopResult(v *AlibabaCloudgameInteractiveGamePlayerStopResult) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	poolAlibabaCloudgameInteractiveGamePlayerStopResult.Put(v)
}
