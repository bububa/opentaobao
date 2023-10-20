package cloudgame

import (
	"sync"
)

// AlibabaCloudgameInteractiveGameStatusGetResult 结构体
type AlibabaCloudgameInteractiveGameStatusGetResult struct {
	// 返回状态码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Data *GameStatusGetResponse `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaCloudgameInteractiveGameStatusGetResult = sync.Pool{
	New: func() any {
		return new(AlibabaCloudgameInteractiveGameStatusGetResult)
	},
}

// GetAlibabaCloudgameInteractiveGameStatusGetResult() 从对象池中获取AlibabaCloudgameInteractiveGameStatusGetResult
func GetAlibabaCloudgameInteractiveGameStatusGetResult() *AlibabaCloudgameInteractiveGameStatusGetResult {
	return poolAlibabaCloudgameInteractiveGameStatusGetResult.Get().(*AlibabaCloudgameInteractiveGameStatusGetResult)
}

// ReleaseAlibabaCloudgameInteractiveGameStatusGetResult 释放AlibabaCloudgameInteractiveGameStatusGetResult
func ReleaseAlibabaCloudgameInteractiveGameStatusGetResult(v *AlibabaCloudgameInteractiveGameStatusGetResult) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	poolAlibabaCloudgameInteractiveGameStatusGetResult.Put(v)
}
