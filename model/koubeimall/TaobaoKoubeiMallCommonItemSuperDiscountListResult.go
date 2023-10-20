package koubeimall

import (
	"sync"
)

// TaobaoKoubeiMallCommonItemSuperDiscountListResult 结构体
type TaobaoKoubeiMallCommonItemSuperDiscountListResult struct {
	// API请求全链路追踪ID
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 超值特惠商品模型
	Data *SuperDiscountDto `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息，success=false时，返回相关错误信息
	Error *TribeError `json:"error,omitempty" xml:"error,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoKoubeiMallCommonItemSuperDiscountListResult = sync.Pool{
	New: func() any {
		return new(TaobaoKoubeiMallCommonItemSuperDiscountListResult)
	},
}

// GetTaobaoKoubeiMallCommonItemSuperDiscountListResult() 从对象池中获取TaobaoKoubeiMallCommonItemSuperDiscountListResult
func GetTaobaoKoubeiMallCommonItemSuperDiscountListResult() *TaobaoKoubeiMallCommonItemSuperDiscountListResult {
	return poolTaobaoKoubeiMallCommonItemSuperDiscountListResult.Get().(*TaobaoKoubeiMallCommonItemSuperDiscountListResult)
}

// ReleaseTaobaoKoubeiMallCommonItemSuperDiscountListResult 释放TaobaoKoubeiMallCommonItemSuperDiscountListResult
func ReleaseTaobaoKoubeiMallCommonItemSuperDiscountListResult(v *TaobaoKoubeiMallCommonItemSuperDiscountListResult) {
	v.TraceId = ""
	v.Data = nil
	v.Error = nil
	v.Success = false
	poolTaobaoKoubeiMallCommonItemSuperDiscountListResult.Put(v)
}
