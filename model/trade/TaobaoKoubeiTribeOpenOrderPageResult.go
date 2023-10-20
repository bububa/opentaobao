package trade

import (
	"sync"
)

// TaobaoKoubeiTribeOpenOrderPageResult 结构体
type TaobaoKoubeiTribeOpenOrderPageResult struct {
	// request唯一ID
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 错误提示
	Error string `json:"error,omitempty" xml:"error,omitempty"`
	// 订单信息结果
	Data *OrderInfoResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoKoubeiTribeOpenOrderPageResult = sync.Pool{
	New: func() any {
		return new(TaobaoKoubeiTribeOpenOrderPageResult)
	},
}

// GetTaobaoKoubeiTribeOpenOrderPageResult() 从对象池中获取TaobaoKoubeiTribeOpenOrderPageResult
func GetTaobaoKoubeiTribeOpenOrderPageResult() *TaobaoKoubeiTribeOpenOrderPageResult {
	return poolTaobaoKoubeiTribeOpenOrderPageResult.Get().(*TaobaoKoubeiTribeOpenOrderPageResult)
}

// ReleaseTaobaoKoubeiTribeOpenOrderPageResult 释放TaobaoKoubeiTribeOpenOrderPageResult
func ReleaseTaobaoKoubeiTribeOpenOrderPageResult(v *TaobaoKoubeiTribeOpenOrderPageResult) {
	v.TraceId = ""
	v.Error = ""
	v.Data = nil
	v.Success = false
	poolTaobaoKoubeiTribeOpenOrderPageResult.Put(v)
}
