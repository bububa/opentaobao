package tmallservice

import (
	"sync"
)

// WorkcardRequestSource 结构体
type WorkcardRequestSource struct {
	// 入驻服务平台淘宝账号
	RealTpNick string `json:"real_tp_nick,omitempty" xml:"real_tp_nick,omitempty"`
	// 门店Code
	ServiceStoreCode string `json:"service_store_code,omitempty" xml:"service_store_code,omitempty"`
}

var poolWorkcardRequestSource = sync.Pool{
	New: func() any {
		return new(WorkcardRequestSource)
	},
}

// GetWorkcardRequestSource() 从对象池中获取WorkcardRequestSource
func GetWorkcardRequestSource() *WorkcardRequestSource {
	return poolWorkcardRequestSource.Get().(*WorkcardRequestSource)
}

// ReleaseWorkcardRequestSource 释放WorkcardRequestSource
func ReleaseWorkcardRequestSource(v *WorkcardRequestSource) {
	v.RealTpNick = ""
	v.ServiceStoreCode = ""
	poolWorkcardRequestSource.Put(v)
}
