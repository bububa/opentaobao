package koubeimall

import (
	"sync"
)

// TaobaoKoubeiMallCommonStorePageResult 结构体
type TaobaoKoubeiMallCommonStorePageResult struct {
	// API请求全链路追踪ID
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 分页模型
	Data *PageResult `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息，success=false时，返回相关错误信息
	Error *TribeError `json:"error,omitempty" xml:"error,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoKoubeiMallCommonStorePageResult = sync.Pool{
	New: func() any {
		return new(TaobaoKoubeiMallCommonStorePageResult)
	},
}

// GetTaobaoKoubeiMallCommonStorePageResult() 从对象池中获取TaobaoKoubeiMallCommonStorePageResult
func GetTaobaoKoubeiMallCommonStorePageResult() *TaobaoKoubeiMallCommonStorePageResult {
	return poolTaobaoKoubeiMallCommonStorePageResult.Get().(*TaobaoKoubeiMallCommonStorePageResult)
}

// ReleaseTaobaoKoubeiMallCommonStorePageResult 释放TaobaoKoubeiMallCommonStorePageResult
func ReleaseTaobaoKoubeiMallCommonStorePageResult(v *TaobaoKoubeiMallCommonStorePageResult) {
	v.TraceId = ""
	v.Data = nil
	v.Error = nil
	v.Success = false
	poolTaobaoKoubeiMallCommonStorePageResult.Put(v)
}
