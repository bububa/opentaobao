package baichuan

import (
	"sync"
)

// TaobaoBaichuanItemsSubscribeResult 结构体
type TaobaoBaichuanItemsSubscribeResult struct {
	// 按不同的返回码将结果分部分返回
	ResultList []ResultMeta `json:"result_list,omitempty" xml:"result_list>result_meta,omitempty"`
}

var poolTaobaoBaichuanItemsSubscribeResult = sync.Pool{
	New: func() any {
		return new(TaobaoBaichuanItemsSubscribeResult)
	},
}

// GetTaobaoBaichuanItemsSubscribeResult() 从对象池中获取TaobaoBaichuanItemsSubscribeResult
func GetTaobaoBaichuanItemsSubscribeResult() *TaobaoBaichuanItemsSubscribeResult {
	return poolTaobaoBaichuanItemsSubscribeResult.Get().(*TaobaoBaichuanItemsSubscribeResult)
}

// ReleaseTaobaoBaichuanItemsSubscribeResult 释放TaobaoBaichuanItemsSubscribeResult
func ReleaseTaobaoBaichuanItemsSubscribeResult(v *TaobaoBaichuanItemsSubscribeResult) {
	v.ResultList = v.ResultList[:0]
	poolTaobaoBaichuanItemsSubscribeResult.Put(v)
}
