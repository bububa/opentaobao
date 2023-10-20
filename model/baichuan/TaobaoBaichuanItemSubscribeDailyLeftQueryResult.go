package baichuan

import (
	"sync"
)

// TaobaoBaichuanItemSubscribeDailyLeftQueryResult 结构体
type TaobaoBaichuanItemSubscribeDailyLeftQueryResult struct {
	// 返回
	ResultList []ResultMeta `json:"result_list,omitempty" xml:"result_list>result_meta,omitempty"`
}

var poolTaobaoBaichuanItemSubscribeDailyLeftQueryResult = sync.Pool{
	New: func() any {
		return new(TaobaoBaichuanItemSubscribeDailyLeftQueryResult)
	},
}

// GetTaobaoBaichuanItemSubscribeDailyLeftQueryResult() 从对象池中获取TaobaoBaichuanItemSubscribeDailyLeftQueryResult
func GetTaobaoBaichuanItemSubscribeDailyLeftQueryResult() *TaobaoBaichuanItemSubscribeDailyLeftQueryResult {
	return poolTaobaoBaichuanItemSubscribeDailyLeftQueryResult.Get().(*TaobaoBaichuanItemSubscribeDailyLeftQueryResult)
}

// ReleaseTaobaoBaichuanItemSubscribeDailyLeftQueryResult 释放TaobaoBaichuanItemSubscribeDailyLeftQueryResult
func ReleaseTaobaoBaichuanItemSubscribeDailyLeftQueryResult(v *TaobaoBaichuanItemSubscribeDailyLeftQueryResult) {
	v.ResultList = v.ResultList[:0]
	poolTaobaoBaichuanItemSubscribeDailyLeftQueryResult.Put(v)
}
