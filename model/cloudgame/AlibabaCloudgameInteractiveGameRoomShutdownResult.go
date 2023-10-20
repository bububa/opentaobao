package cloudgame

import (
	"sync"
)

// AlibabaCloudgameInteractiveGameRoomShutdownResult 结构体
type AlibabaCloudgameInteractiveGameRoomShutdownResult struct {
	// 状态码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 强制停止房间返回结果
	Data *ShutdownRoomResponse `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaCloudgameInteractiveGameRoomShutdownResult = sync.Pool{
	New: func() any {
		return new(AlibabaCloudgameInteractiveGameRoomShutdownResult)
	},
}

// GetAlibabaCloudgameInteractiveGameRoomShutdownResult() 从对象池中获取AlibabaCloudgameInteractiveGameRoomShutdownResult
func GetAlibabaCloudgameInteractiveGameRoomShutdownResult() *AlibabaCloudgameInteractiveGameRoomShutdownResult {
	return poolAlibabaCloudgameInteractiveGameRoomShutdownResult.Get().(*AlibabaCloudgameInteractiveGameRoomShutdownResult)
}

// ReleaseAlibabaCloudgameInteractiveGameRoomShutdownResult 释放AlibabaCloudgameInteractiveGameRoomShutdownResult
func ReleaseAlibabaCloudgameInteractiveGameRoomShutdownResult(v *AlibabaCloudgameInteractiveGameRoomShutdownResult) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	poolAlibabaCloudgameInteractiveGameRoomShutdownResult.Put(v)
}
