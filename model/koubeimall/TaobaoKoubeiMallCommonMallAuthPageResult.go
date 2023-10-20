package koubeimall

import (
	"sync"
)

// TaobaoKoubeiMallCommonMallAuthPageResult 结构体
type TaobaoKoubeiMallCommonMallAuthPageResult struct {
	// API请求全链路追踪ID
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// API返回的data模型
	Data *PageResult `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息，success=false时，返回相关错误信息
	Error *TribeError `json:"error,omitempty" xml:"error,omitempty"`
	// 请求是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoKoubeiMallCommonMallAuthPageResult = sync.Pool{
	New: func() any {
		return new(TaobaoKoubeiMallCommonMallAuthPageResult)
	},
}

// GetTaobaoKoubeiMallCommonMallAuthPageResult() 从对象池中获取TaobaoKoubeiMallCommonMallAuthPageResult
func GetTaobaoKoubeiMallCommonMallAuthPageResult() *TaobaoKoubeiMallCommonMallAuthPageResult {
	return poolTaobaoKoubeiMallCommonMallAuthPageResult.Get().(*TaobaoKoubeiMallCommonMallAuthPageResult)
}

// ReleaseTaobaoKoubeiMallCommonMallAuthPageResult 释放TaobaoKoubeiMallCommonMallAuthPageResult
func ReleaseTaobaoKoubeiMallCommonMallAuthPageResult(v *TaobaoKoubeiMallCommonMallAuthPageResult) {
	v.TraceId = ""
	v.Data = nil
	v.Error = nil
	v.Success = false
	poolTaobaoKoubeiMallCommonMallAuthPageResult.Put(v)
}
