package logistic

import (
	"sync"
)

// Addressdtos 结构体
type Addressdtos struct {
	// sender address
	Sender *AeopWlDeclareAddressDto `json:"sender,omitempty" xml:"sender,omitempty"`
	// pickup address
	Pickup *AeopWlDeclareAddressDto `json:"pickup,omitempty" xml:"pickup,omitempty"`
	// receiver address
	Receiver *AeopWlDeclareAddressDto `json:"receiver,omitempty" xml:"receiver,omitempty"`
	// refund address
	Refund *AeopWlDeclareAddressDto `json:"refund,omitempty" xml:"refund,omitempty"`
}

var poolAddressdtos = sync.Pool{
	New: func() any {
		return new(Addressdtos)
	},
}

// GetAddressdtos() 从对象池中获取Addressdtos
func GetAddressdtos() *Addressdtos {
	return poolAddressdtos.Get().(*Addressdtos)
}

// ReleaseAddressdtos 释放Addressdtos
func ReleaseAddressdtos(v *Addressdtos) {
	v.Sender = nil
	v.Pickup = nil
	v.Receiver = nil
	v.Refund = nil
	poolAddressdtos.Put(v)
}
