package cloudgame

import (
	"sync"
)

// AlibabaCloudgameInteractiveGamePlayerStatusGetResult 结构体
type AlibabaCloudgameInteractiveGamePlayerStatusGetResult struct {
	// 返回状态码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Data *GameStatusGetResponse `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaCloudgameInteractiveGamePlayerStatusGetResult = sync.Pool{
	New: func() any {
		return new(AlibabaCloudgameInteractiveGamePlayerStatusGetResult)
	},
}

// GetAlibabaCloudgameInteractiveGamePlayerStatusGetResult() 从对象池中获取AlibabaCloudgameInteractiveGamePlayerStatusGetResult
func GetAlibabaCloudgameInteractiveGamePlayerStatusGetResult() *AlibabaCloudgameInteractiveGamePlayerStatusGetResult {
	return poolAlibabaCloudgameInteractiveGamePlayerStatusGetResult.Get().(*AlibabaCloudgameInteractiveGamePlayerStatusGetResult)
}

// ReleaseAlibabaCloudgameInteractiveGamePlayerStatusGetResult 释放AlibabaCloudgameInteractiveGamePlayerStatusGetResult
func ReleaseAlibabaCloudgameInteractiveGamePlayerStatusGetResult(v *AlibabaCloudgameInteractiveGamePlayerStatusGetResult) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	poolAlibabaCloudgameInteractiveGamePlayerStatusGetResult.Put(v)
}
