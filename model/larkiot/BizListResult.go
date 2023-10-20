package larkiot

import (
	"sync"
)

// BizListResult 结构体
type BizListResult struct {
	// 影院列表
	DataList []CinemaIotResponse `json:"data_list,omitempty" xml:"data_list>cinema_iot_response,omitempty"`
	// 业务码
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 业务信息
	BizMsg string `json:"biz_msg,omitempty" xml:"biz_msg,omitempty"`
	// 访问结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolBizListResult = sync.Pool{
	New: func() any {
		return new(BizListResult)
	},
}

// GetBizListResult() 从对象池中获取BizListResult
func GetBizListResult() *BizListResult {
	return poolBizListResult.Get().(*BizListResult)
}

// ReleaseBizListResult 释放BizListResult
func ReleaseBizListResult(v *BizListResult) {
	v.DataList = v.DataList[:0]
	v.BizCode = ""
	v.BizMsg = ""
	v.Success = false
	poolBizListResult.Put(v)
}
