package wdk

import (
	"sync"
)

// FulfillLogisticListResult 结构体
type FulfillLogisticListResult struct {
	// 小票批次信息
	Results []ReceiptBatchInfo `json:"results,omitempty" xml:"results>receipt_batch_info,omitempty"`
	// SYSTEM_ERROR(&#34;SYSTEM_ERROR&#34;, &#34;系统异常&#34;),PARAM_ERROR(&#34;PARAM_ERROR&#34;, &#34;参数错误&#34;),BUSINESS_ERROR(&#34;BUSINESS_ERROR&#34;, &#34;业务异常&#34;);
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 返回码含义描述
	ErrDesc string `json:"err_desc,omitempty" xml:"err_desc,omitempty"`
	// true 调用成功 false 调用异常
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolFulfillLogisticListResult = sync.Pool{
	New: func() any {
		return new(FulfillLogisticListResult)
	},
}

// GetFulfillLogisticListResult() 从对象池中获取FulfillLogisticListResult
func GetFulfillLogisticListResult() *FulfillLogisticListResult {
	return poolFulfillLogisticListResult.Get().(*FulfillLogisticListResult)
}

// ReleaseFulfillLogisticListResult 释放FulfillLogisticListResult
func ReleaseFulfillLogisticListResult(v *FulfillLogisticListResult) {
	v.Results = v.Results[:0]
	v.ErrCode = ""
	v.ErrDesc = ""
	v.Success = false
	poolFulfillLogisticListResult.Put(v)
}
