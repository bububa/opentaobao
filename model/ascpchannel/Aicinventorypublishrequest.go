package ascpchannel

import (
	"sync"
)

// Aicinventorypublishrequest 结构体
type Aicinventorypublishrequest struct {
	// 库存发布请求参数
	InventoryMainOperation []Inventorymainoperation `json:"inventory_main_operation,omitempty" xml:"inventory_main_operation>inventorymainoperation,omitempty"`
}

var poolAicinventorypublishrequest = sync.Pool{
	New: func() any {
		return new(Aicinventorypublishrequest)
	},
}

// GetAicinventorypublishrequest() 从对象池中获取Aicinventorypublishrequest
func GetAicinventorypublishrequest() *Aicinventorypublishrequest {
	return poolAicinventorypublishrequest.Get().(*Aicinventorypublishrequest)
}

// ReleaseAicinventorypublishrequest 释放Aicinventorypublishrequest
func ReleaseAicinventorypublishrequest(v *Aicinventorypublishrequest) {
	v.InventoryMainOperation = v.InventoryMainOperation[:0]
	poolAicinventorypublishrequest.Put(v)
}
