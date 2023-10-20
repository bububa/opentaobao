package ieagency

import (
	"sync"
)

// BookPayOrderRq 结构体
type BookPayOrderRq struct {
	// 乘机人信息
	PassengerParams []PassengerParam `json:"passenger_params,omitempty" xml:"passenger_params>passenger_param,omitempty"`
	// 代理商飞猪昵称（注意不是淘宝代理商昵称）
	AgentName string `json:"agent_name,omitempty" xml:"agent_name,omitempty"`
	// 外部采购订单标识（有唯一性要求）
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 联系人信息
	ContactsParam *ContactsParam `json:"contacts_param,omitempty" xml:"contacts_param,omitempty"`
	// 航班产品信息
	ItemParam *ItemParam `json:"item_param,omitempty" xml:"item_param,omitempty"`
}

var poolBookPayOrderRq = sync.Pool{
	New: func() any {
		return new(BookPayOrderRq)
	},
}

// GetBookPayOrderRq() 从对象池中获取BookPayOrderRq
func GetBookPayOrderRq() *BookPayOrderRq {
	return poolBookPayOrderRq.Get().(*BookPayOrderRq)
}

// ReleaseBookPayOrderRq 释放BookPayOrderRq
func ReleaseBookPayOrderRq(v *BookPayOrderRq) {
	v.PassengerParams = v.PassengerParams[:0]
	v.AgentName = ""
	v.OutOrderId = ""
	v.ContactsParam = nil
	v.ItemParam = nil
	poolBookPayOrderRq.Put(v)
}
