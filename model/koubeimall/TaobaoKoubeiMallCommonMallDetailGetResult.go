package koubeimall

import (
	"sync"
)

// TaobaoKoubeiMallCommonMallDetailGetResult 结构体
type TaobaoKoubeiMallCommonMallDetailGetResult struct {
	// API请求全链路追踪ID
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// API返回的商圈data模型
	Data *MallDetailDto `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息，success=false时，返回相关错误信息
	Error *TribeError `json:"error,omitempty" xml:"error,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoKoubeiMallCommonMallDetailGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoKoubeiMallCommonMallDetailGetResult)
	},
}

// GetTaobaoKoubeiMallCommonMallDetailGetResult() 从对象池中获取TaobaoKoubeiMallCommonMallDetailGetResult
func GetTaobaoKoubeiMallCommonMallDetailGetResult() *TaobaoKoubeiMallCommonMallDetailGetResult {
	return poolTaobaoKoubeiMallCommonMallDetailGetResult.Get().(*TaobaoKoubeiMallCommonMallDetailGetResult)
}

// ReleaseTaobaoKoubeiMallCommonMallDetailGetResult 释放TaobaoKoubeiMallCommonMallDetailGetResult
func ReleaseTaobaoKoubeiMallCommonMallDetailGetResult(v *TaobaoKoubeiMallCommonMallDetailGetResult) {
	v.TraceId = ""
	v.Data = nil
	v.Error = nil
	v.Success = false
	poolTaobaoKoubeiMallCommonMallDetailGetResult.Put(v)
}
