package koubeimall

import (
	"sync"
)

// TaobaoKoubeiMallCommonItemDetailQueryResult 结构体
type TaobaoKoubeiMallCommonItemDetailQueryResult struct {
	// API请求全链路追踪ID
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 商品详情返回结果模型
	Data *ItemDetailResult `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息，success=false时，返回相关错误信息
	Error *TribeError `json:"error,omitempty" xml:"error,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoKoubeiMallCommonItemDetailQueryResult = sync.Pool{
	New: func() any {
		return new(TaobaoKoubeiMallCommonItemDetailQueryResult)
	},
}

// GetTaobaoKoubeiMallCommonItemDetailQueryResult() 从对象池中获取TaobaoKoubeiMallCommonItemDetailQueryResult
func GetTaobaoKoubeiMallCommonItemDetailQueryResult() *TaobaoKoubeiMallCommonItemDetailQueryResult {
	return poolTaobaoKoubeiMallCommonItemDetailQueryResult.Get().(*TaobaoKoubeiMallCommonItemDetailQueryResult)
}

// ReleaseTaobaoKoubeiMallCommonItemDetailQueryResult 释放TaobaoKoubeiMallCommonItemDetailQueryResult
func ReleaseTaobaoKoubeiMallCommonItemDetailQueryResult(v *TaobaoKoubeiMallCommonItemDetailQueryResult) {
	v.TraceId = ""
	v.Data = nil
	v.Error = nil
	v.Success = false
	poolTaobaoKoubeiMallCommonItemDetailQueryResult.Put(v)
}
