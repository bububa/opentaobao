package drugtrace

import (
	"sync"
)

// ProduceInfoList 结构体
type ProduceInfoList struct {
	// 有效期至
	ExpireDate string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
}

var poolProduceInfoList = sync.Pool{
	New: func() any {
		return new(ProduceInfoList)
	},
}

// GetProduceInfoList() 从对象池中获取ProduceInfoList
func GetProduceInfoList() *ProduceInfoList {
	return poolProduceInfoList.Get().(*ProduceInfoList)
}

// ReleaseProduceInfoList 释放ProduceInfoList
func ReleaseProduceInfoList(v *ProduceInfoList) {
	v.ExpireDate = ""
	poolProduceInfoList.Put(v)
}
