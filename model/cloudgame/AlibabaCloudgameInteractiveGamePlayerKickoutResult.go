package cloudgame

import (
	"sync"
)

// AlibabaCloudgameInteractiveGamePlayerKickoutResult 结构体
type AlibabaCloudgameInteractiveGamePlayerKickoutResult struct {
	// 返回状态码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Data *JoinCodeAssignResponse `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaCloudgameInteractiveGamePlayerKickoutResult = sync.Pool{
	New: func() any {
		return new(AlibabaCloudgameInteractiveGamePlayerKickoutResult)
	},
}

// GetAlibabaCloudgameInteractiveGamePlayerKickoutResult() 从对象池中获取AlibabaCloudgameInteractiveGamePlayerKickoutResult
func GetAlibabaCloudgameInteractiveGamePlayerKickoutResult() *AlibabaCloudgameInteractiveGamePlayerKickoutResult {
	return poolAlibabaCloudgameInteractiveGamePlayerKickoutResult.Get().(*AlibabaCloudgameInteractiveGamePlayerKickoutResult)
}

// ReleaseAlibabaCloudgameInteractiveGamePlayerKickoutResult 释放AlibabaCloudgameInteractiveGamePlayerKickoutResult
func ReleaseAlibabaCloudgameInteractiveGamePlayerKickoutResult(v *AlibabaCloudgameInteractiveGamePlayerKickoutResult) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	poolAlibabaCloudgameInteractiveGamePlayerKickoutResult.Put(v)
}
