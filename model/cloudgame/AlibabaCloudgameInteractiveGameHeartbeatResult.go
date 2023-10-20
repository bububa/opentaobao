package cloudgame

import (
	"sync"
)

// AlibabaCloudgameInteractiveGameHeartbeatResult 结构体
type AlibabaCloudgameInteractiveGameHeartbeatResult struct {
	// 返回状态码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Data *HeartBeatResponse `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaCloudgameInteractiveGameHeartbeatResult = sync.Pool{
	New: func() any {
		return new(AlibabaCloudgameInteractiveGameHeartbeatResult)
	},
}

// GetAlibabaCloudgameInteractiveGameHeartbeatResult() 从对象池中获取AlibabaCloudgameInteractiveGameHeartbeatResult
func GetAlibabaCloudgameInteractiveGameHeartbeatResult() *AlibabaCloudgameInteractiveGameHeartbeatResult {
	return poolAlibabaCloudgameInteractiveGameHeartbeatResult.Get().(*AlibabaCloudgameInteractiveGameHeartbeatResult)
}

// ReleaseAlibabaCloudgameInteractiveGameHeartbeatResult 释放AlibabaCloudgameInteractiveGameHeartbeatResult
func ReleaseAlibabaCloudgameInteractiveGameHeartbeatResult(v *AlibabaCloudgameInteractiveGameHeartbeatResult) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	poolAlibabaCloudgameInteractiveGameHeartbeatResult.Put(v)
}
