package baichuan

import (
	"sync"
)

// TaobaoBaichuanItemSubscribeRelationQueryResult 结构体
type TaobaoBaichuanItemSubscribeRelationQueryResult struct {
	// 返回的list
	ResultList []ResultMeta `json:"result_list,omitempty" xml:"result_list>result_meta,omitempty"`
}

var poolTaobaoBaichuanItemSubscribeRelationQueryResult = sync.Pool{
	New: func() any {
		return new(TaobaoBaichuanItemSubscribeRelationQueryResult)
	},
}

// GetTaobaoBaichuanItemSubscribeRelationQueryResult() 从对象池中获取TaobaoBaichuanItemSubscribeRelationQueryResult
func GetTaobaoBaichuanItemSubscribeRelationQueryResult() *TaobaoBaichuanItemSubscribeRelationQueryResult {
	return poolTaobaoBaichuanItemSubscribeRelationQueryResult.Get().(*TaobaoBaichuanItemSubscribeRelationQueryResult)
}

// ReleaseTaobaoBaichuanItemSubscribeRelationQueryResult 释放TaobaoBaichuanItemSubscribeRelationQueryResult
func ReleaseTaobaoBaichuanItemSubscribeRelationQueryResult(v *TaobaoBaichuanItemSubscribeRelationQueryResult) {
	v.ResultList = v.ResultList[:0]
	poolTaobaoBaichuanItemSubscribeRelationQueryResult.Put(v)
}
