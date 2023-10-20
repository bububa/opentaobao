package ascpchannel

import (
	"sync"
)

// Presalesorderconsignconfirmrequest 结构体
type Presalesorderconsignconfirmrequest struct {
	// 预售订单信息
	PresalesOrder *Presalesorder `json:"presales_order,omitempty" xml:"presales_order,omitempty"`
}

var poolPresalesorderconsignconfirmrequest = sync.Pool{
	New: func() any {
		return new(Presalesorderconsignconfirmrequest)
	},
}

// GetPresalesorderconsignconfirmrequest() 从对象池中获取Presalesorderconsignconfirmrequest
func GetPresalesorderconsignconfirmrequest() *Presalesorderconsignconfirmrequest {
	return poolPresalesorderconsignconfirmrequest.Get().(*Presalesorderconsignconfirmrequest)
}

// ReleasePresalesorderconsignconfirmrequest 释放Presalesorderconsignconfirmrequest
func ReleasePresalesorderconsignconfirmrequest(v *Presalesorderconsignconfirmrequest) {
	v.PresalesOrder = nil
	poolPresalesorderconsignconfirmrequest.Put(v)
}
