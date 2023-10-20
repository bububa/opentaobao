package usergrowth

import (
	"sync"
)

// BatchAskResultV2 结构体
type BatchAskResultV2 struct {
	// 匹配的设备与其任务信息列表
	Results []BatchAskResultItem `json:"results,omitempty" xml:"results>batch_ask_result_item,omitempty"`
	// 错误码， 0： 成功；1：限流；2：服务不可用
	Errcode int64 `json:"errcode,omitempty" xml:"errcode,omitempty"`
}

var poolBatchAskResultV2 = sync.Pool{
	New: func() any {
		return new(BatchAskResultV2)
	},
}

// GetBatchAskResultV2() 从对象池中获取BatchAskResultV2
func GetBatchAskResultV2() *BatchAskResultV2 {
	return poolBatchAskResultV2.Get().(*BatchAskResultV2)
}

// ReleaseBatchAskResultV2 释放BatchAskResultV2
func ReleaseBatchAskResultV2(v *BatchAskResultV2) {
	v.Results = v.Results[:0]
	v.Errcode = 0
	poolBatchAskResultV2.Put(v)
}
