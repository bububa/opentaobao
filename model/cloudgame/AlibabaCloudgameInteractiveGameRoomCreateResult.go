package cloudgame

import (
	"sync"
)

// AlibabaCloudgameInteractiveGameRoomCreateResult 结构体
type AlibabaCloudgameInteractiveGameRoomCreateResult struct {
	// 状态码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回提示信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Data *CreateRoomResponse `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaCloudgameInteractiveGameRoomCreateResult = sync.Pool{
	New: func() any {
		return new(AlibabaCloudgameInteractiveGameRoomCreateResult)
	},
}

// GetAlibabaCloudgameInteractiveGameRoomCreateResult() 从对象池中获取AlibabaCloudgameInteractiveGameRoomCreateResult
func GetAlibabaCloudgameInteractiveGameRoomCreateResult() *AlibabaCloudgameInteractiveGameRoomCreateResult {
	return poolAlibabaCloudgameInteractiveGameRoomCreateResult.Get().(*AlibabaCloudgameInteractiveGameRoomCreateResult)
}

// ReleaseAlibabaCloudgameInteractiveGameRoomCreateResult 释放AlibabaCloudgameInteractiveGameRoomCreateResult
func ReleaseAlibabaCloudgameInteractiveGameRoomCreateResult(v *AlibabaCloudgameInteractiveGameRoomCreateResult) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	poolAlibabaCloudgameInteractiveGameRoomCreateResult.Put(v)
}
