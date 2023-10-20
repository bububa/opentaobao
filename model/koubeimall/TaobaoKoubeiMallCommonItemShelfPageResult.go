package koubeimall

import (
	"sync"
)

// TaobaoKoubeiMallCommonItemShelfPageResult 结构体
type TaobaoKoubeiMallCommonItemShelfPageResult struct {
	// API请求全链路追踪ID
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 门店货架商品列表信息
	Data *ShopShelfDto `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息，success=false时，返回相关错误信息
	Error *TribeError `json:"error,omitempty" xml:"error,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoKoubeiMallCommonItemShelfPageResult = sync.Pool{
	New: func() any {
		return new(TaobaoKoubeiMallCommonItemShelfPageResult)
	},
}

// GetTaobaoKoubeiMallCommonItemShelfPageResult() 从对象池中获取TaobaoKoubeiMallCommonItemShelfPageResult
func GetTaobaoKoubeiMallCommonItemShelfPageResult() *TaobaoKoubeiMallCommonItemShelfPageResult {
	return poolTaobaoKoubeiMallCommonItemShelfPageResult.Get().(*TaobaoKoubeiMallCommonItemShelfPageResult)
}

// ReleaseTaobaoKoubeiMallCommonItemShelfPageResult 释放TaobaoKoubeiMallCommonItemShelfPageResult
func ReleaseTaobaoKoubeiMallCommonItemShelfPageResult(v *TaobaoKoubeiMallCommonItemShelfPageResult) {
	v.TraceId = ""
	v.Data = nil
	v.Error = nil
	v.Success = false
	poolTaobaoKoubeiMallCommonItemShelfPageResult.Put(v)
}
