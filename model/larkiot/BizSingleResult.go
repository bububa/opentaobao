package larkiot

import (
	"sync"
)

// BizSingleResult 结构体
type BizSingleResult struct {
	// 业务订结果
	BizMsg string `json:"biz_msg,omitempty" xml:"biz_msg,omitempty"`
	// 业务码
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 业务返回
	Data *IotGoodsOrderRsp `json:"data,omitempty" xml:"data,omitempty"`
	// 接口是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolBizSingleResult = sync.Pool{
	New: func() any {
		return new(BizSingleResult)
	},
}

// GetBizSingleResult() 从对象池中获取BizSingleResult
func GetBizSingleResult() *BizSingleResult {
	return poolBizSingleResult.Get().(*BizSingleResult)
}

// ReleaseBizSingleResult 释放BizSingleResult
func ReleaseBizSingleResult(v *BizSingleResult) {
	v.BizMsg = ""
	v.BizCode = ""
	v.Data = nil
	v.Success = false
	poolBizSingleResult.Put(v)
}
