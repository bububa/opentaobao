package baichuan

import (
	"sync"
)

// TaobaoBaichuanItemSubscribeResult 结构体
type TaobaoBaichuanItemSubscribeResult struct {
	// 返回的列表
	ResultList []ResultMeta `json:"result_list,omitempty" xml:"result_list>result_meta,omitempty"`
}

var poolTaobaoBaichuanItemSubscribeResult = sync.Pool{
	New: func() any {
		return new(TaobaoBaichuanItemSubscribeResult)
	},
}

// GetTaobaoBaichuanItemSubscribeResult() 从对象池中获取TaobaoBaichuanItemSubscribeResult
func GetTaobaoBaichuanItemSubscribeResult() *TaobaoBaichuanItemSubscribeResult {
	return poolTaobaoBaichuanItemSubscribeResult.Get().(*TaobaoBaichuanItemSubscribeResult)
}

// ReleaseTaobaoBaichuanItemSubscribeResult 释放TaobaoBaichuanItemSubscribeResult
func ReleaseTaobaoBaichuanItemSubscribeResult(v *TaobaoBaichuanItemSubscribeResult) {
	v.ResultList = v.ResultList[:0]
	poolTaobaoBaichuanItemSubscribeResult.Put(v)
}
