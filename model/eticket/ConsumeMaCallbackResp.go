package eticket

import (
	"sync"
)

// ConsumeMaCallbackResp 结构体
type ConsumeMaCallbackResp struct {
	// 业务回复KV
	AttributeMap string `json:"attribute_map,omitempty" xml:"attribute_map,omitempty"`
}

var poolConsumeMaCallbackResp = sync.Pool{
	New: func() any {
		return new(ConsumeMaCallbackResp)
	},
}

// GetConsumeMaCallbackResp() 从对象池中获取ConsumeMaCallbackResp
func GetConsumeMaCallbackResp() *ConsumeMaCallbackResp {
	return poolConsumeMaCallbackResp.Get().(*ConsumeMaCallbackResp)
}

// ReleaseConsumeMaCallbackResp 释放ConsumeMaCallbackResp
func ReleaseConsumeMaCallbackResp(v *ConsumeMaCallbackResp) {
	v.AttributeMap = ""
	poolConsumeMaCallbackResp.Put(v)
}
