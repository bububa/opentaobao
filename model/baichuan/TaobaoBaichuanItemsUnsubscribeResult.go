package baichuan

import (
	"sync"
)

// TaobaoBaichuanItemsUnsubscribeResult 结构体
type TaobaoBaichuanItemsUnsubscribeResult struct {
	// 返回按resultCode分为多个返回部分
	ResultList []ResultMeta `json:"result_list,omitempty" xml:"result_list>result_meta,omitempty"`
}

var poolTaobaoBaichuanItemsUnsubscribeResult = sync.Pool{
	New: func() any {
		return new(TaobaoBaichuanItemsUnsubscribeResult)
	},
}

// GetTaobaoBaichuanItemsUnsubscribeResult() 从对象池中获取TaobaoBaichuanItemsUnsubscribeResult
func GetTaobaoBaichuanItemsUnsubscribeResult() *TaobaoBaichuanItemsUnsubscribeResult {
	return poolTaobaoBaichuanItemsUnsubscribeResult.Get().(*TaobaoBaichuanItemsUnsubscribeResult)
}

// ReleaseTaobaoBaichuanItemsUnsubscribeResult 释放TaobaoBaichuanItemsUnsubscribeResult
func ReleaseTaobaoBaichuanItemsUnsubscribeResult(v *TaobaoBaichuanItemsUnsubscribeResult) {
	v.ResultList = v.ResultList[:0]
	poolTaobaoBaichuanItemsUnsubscribeResult.Put(v)
}
