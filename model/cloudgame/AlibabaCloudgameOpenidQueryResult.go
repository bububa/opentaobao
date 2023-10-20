package cloudgame

import (
	"sync"
)

// AlibabaCloudgameOpenidQueryResult 结构体
type AlibabaCloudgameOpenidQueryResult struct {
	// 结果码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 结果码描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回数据
	Data *HavanaUserIdQueryResponseVo `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaCloudgameOpenidQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaCloudgameOpenidQueryResult)
	},
}

// GetAlibabaCloudgameOpenidQueryResult() 从对象池中获取AlibabaCloudgameOpenidQueryResult
func GetAlibabaCloudgameOpenidQueryResult() *AlibabaCloudgameOpenidQueryResult {
	return poolAlibabaCloudgameOpenidQueryResult.Get().(*AlibabaCloudgameOpenidQueryResult)
}

// ReleaseAlibabaCloudgameOpenidQueryResult 释放AlibabaCloudgameOpenidQueryResult
func ReleaseAlibabaCloudgameOpenidQueryResult(v *AlibabaCloudgameOpenidQueryResult) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	poolAlibabaCloudgameOpenidQueryResult.Put(v)
}
