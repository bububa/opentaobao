package tbk

import (
	"sync"
)

// TaobaoTbkDgCpaActivityReportRpcResult 结构体
type TaobaoTbkDgCpaActivityReportRpcResult struct {
	// 分页模型
	Data *PageResult `json:"data,omitempty" xml:"data,omitempty"`
}

var poolTaobaoTbkDgCpaActivityReportRpcResult = sync.Pool{
	New: func() any {
		return new(TaobaoTbkDgCpaActivityReportRpcResult)
	},
}

// GetTaobaoTbkDgCpaActivityReportRpcResult() 从对象池中获取TaobaoTbkDgCpaActivityReportRpcResult
func GetTaobaoTbkDgCpaActivityReportRpcResult() *TaobaoTbkDgCpaActivityReportRpcResult {
	return poolTaobaoTbkDgCpaActivityReportRpcResult.Get().(*TaobaoTbkDgCpaActivityReportRpcResult)
}

// ReleaseTaobaoTbkDgCpaActivityReportRpcResult 释放TaobaoTbkDgCpaActivityReportRpcResult
func ReleaseTaobaoTbkDgCpaActivityReportRpcResult(v *TaobaoTbkDgCpaActivityReportRpcResult) {
	v.Data = nil
	poolTaobaoTbkDgCpaActivityReportRpcResult.Put(v)
}
