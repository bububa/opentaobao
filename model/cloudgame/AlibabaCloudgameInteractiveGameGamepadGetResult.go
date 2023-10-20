package cloudgame

import (
	"sync"
)

// AlibabaCloudgameInteractiveGameGamepadGetResult 结构体
type AlibabaCloudgameInteractiveGameGamepadGetResult struct {
	// 返回码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回数据
	Data *GamepadGetResponse `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaCloudgameInteractiveGameGamepadGetResult = sync.Pool{
	New: func() any {
		return new(AlibabaCloudgameInteractiveGameGamepadGetResult)
	},
}

// GetAlibabaCloudgameInteractiveGameGamepadGetResult() 从对象池中获取AlibabaCloudgameInteractiveGameGamepadGetResult
func GetAlibabaCloudgameInteractiveGameGamepadGetResult() *AlibabaCloudgameInteractiveGameGamepadGetResult {
	return poolAlibabaCloudgameInteractiveGameGamepadGetResult.Get().(*AlibabaCloudgameInteractiveGameGamepadGetResult)
}

// ReleaseAlibabaCloudgameInteractiveGameGamepadGetResult 释放AlibabaCloudgameInteractiveGameGamepadGetResult
func ReleaseAlibabaCloudgameInteractiveGameGamepadGetResult(v *AlibabaCloudgameInteractiveGameGamepadGetResult) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	poolAlibabaCloudgameInteractiveGameGamepadGetResult.Put(v)
}
