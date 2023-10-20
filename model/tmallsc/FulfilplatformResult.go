package tmallsc

import (
	"sync"
)

// FulfilplatformResult 结构体
type FulfilplatformResult struct {
	// 错误描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 每页数据信息
	ResultData *PagedResult `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// true：成功；false：失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolFulfilplatformResult = sync.Pool{
	New: func() any {
		return new(FulfilplatformResult)
	},
}

// GetFulfilplatformResult() 从对象池中获取FulfilplatformResult
func GetFulfilplatformResult() *FulfilplatformResult {
	return poolFulfilplatformResult.Get().(*FulfilplatformResult)
}

// ReleaseFulfilplatformResult 释放FulfilplatformResult
func ReleaseFulfilplatformResult(v *FulfilplatformResult) {
	v.MsgInfo = ""
	v.MsgCode = ""
	v.ResultData = nil
	v.Success = false
	poolFulfilplatformResult.Put(v)
}
