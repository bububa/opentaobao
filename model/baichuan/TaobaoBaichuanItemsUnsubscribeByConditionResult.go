package baichuan

import (
	"sync"
)

// TaobaoBaichuanItemsUnsubscribeByConditionResult 结构体
type TaobaoBaichuanItemsUnsubscribeByConditionResult struct {
	// 分返回码返回结果
	ResultList []ResultMeta `json:"result_list,omitempty" xml:"result_list>result_meta,omitempty"`
}

var poolTaobaoBaichuanItemsUnsubscribeByConditionResult = sync.Pool{
	New: func() any {
		return new(TaobaoBaichuanItemsUnsubscribeByConditionResult)
	},
}

// GetTaobaoBaichuanItemsUnsubscribeByConditionResult() 从对象池中获取TaobaoBaichuanItemsUnsubscribeByConditionResult
func GetTaobaoBaichuanItemsUnsubscribeByConditionResult() *TaobaoBaichuanItemsUnsubscribeByConditionResult {
	return poolTaobaoBaichuanItemsUnsubscribeByConditionResult.Get().(*TaobaoBaichuanItemsUnsubscribeByConditionResult)
}

// ReleaseTaobaoBaichuanItemsUnsubscribeByConditionResult 释放TaobaoBaichuanItemsUnsubscribeByConditionResult
func ReleaseTaobaoBaichuanItemsUnsubscribeByConditionResult(v *TaobaoBaichuanItemsUnsubscribeByConditionResult) {
	v.ResultList = v.ResultList[:0]
	poolTaobaoBaichuanItemsUnsubscribeByConditionResult.Put(v)
}
