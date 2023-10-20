package user

import (
	"sync"
)

// TaobaoKoubeiTribeOpenUserQueryResult 结构体
type TaobaoKoubeiTribeOpenUserQueryResult struct {
	// traceId
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 错误信息
	Error string `json:"error,omitempty" xml:"error,omitempty"`
	// 用户信息DTO
	Data *UserInfoDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoKoubeiTribeOpenUserQueryResult = sync.Pool{
	New: func() any {
		return new(TaobaoKoubeiTribeOpenUserQueryResult)
	},
}

// GetTaobaoKoubeiTribeOpenUserQueryResult() 从对象池中获取TaobaoKoubeiTribeOpenUserQueryResult
func GetTaobaoKoubeiTribeOpenUserQueryResult() *TaobaoKoubeiTribeOpenUserQueryResult {
	return poolTaobaoKoubeiTribeOpenUserQueryResult.Get().(*TaobaoKoubeiTribeOpenUserQueryResult)
}

// ReleaseTaobaoKoubeiTribeOpenUserQueryResult 释放TaobaoKoubeiTribeOpenUserQueryResult
func ReleaseTaobaoKoubeiTribeOpenUserQueryResult(v *TaobaoKoubeiTribeOpenUserQueryResult) {
	v.TraceId = ""
	v.Error = ""
	v.Data = nil
	v.Success = false
	poolTaobaoKoubeiTribeOpenUserQueryResult.Put(v)
}
