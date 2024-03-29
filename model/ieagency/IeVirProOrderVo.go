package ieagency

import (
	"sync"
)

// IeVirProOrderVo 结构体
type IeVirProOrderVo struct {
	// 辅营购买信息
	PassengerAuxVos []IePassengerAuxVo `json:"passenger_aux_vos,omitempty" xml:"passenger_aux_vos>ie_passenger_aux_vo,omitempty"`
	// 辅营订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolIeVirProOrderVo = sync.Pool{
	New: func() any {
		return new(IeVirProOrderVo)
	},
}

// GetIeVirProOrderVo() 从对象池中获取IeVirProOrderVo
func GetIeVirProOrderVo() *IeVirProOrderVo {
	return poolIeVirProOrderVo.Get().(*IeVirProOrderVo)
}

// ReleaseIeVirProOrderVo 释放IeVirProOrderVo
func ReleaseIeVirProOrderVo(v *IeVirProOrderVo) {
	v.PassengerAuxVos = v.PassengerAuxVos[:0]
	v.OrderId = 0
	poolIeVirProOrderVo.Put(v)
}
