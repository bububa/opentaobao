package wdk

import (
	"sync"
)

// EleOrderInfoBo 结构体
type EleOrderInfoBo struct {
	// 损失承担方
	ResponsibleParty string `json:"responsible_party,omitempty" xml:"responsible_party,omitempty"`
	// 金额
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 账务时间
	TradeCreateTime string `json:"trade_create_time,omitempty" xml:"trade_create_time,omitempty"`
	// 下单时间
	OrderCreateTime string `json:"order_create_time,omitempty" xml:"order_create_time,omitempty"`
	// 实际付款主体
	PayEntity string `json:"pay_entity,omitempty" xml:"pay_entity,omitempty"`
	// 饿了么订单id
	EleOrderId string `json:"ele_order_id,omitempty" xml:"ele_order_id,omitempty"`
	// 订单id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 订单来源
	OrderFrom string `json:"order_from,omitempty" xml:"order_from,omitempty"`
	// 订单序号
	OrderIndex string `json:"order_index,omitempty" xml:"order_index,omitempty"`
	// 订单费用明细
	OrderDetailFee *OrderDetailFee `json:"order_detail_fee,omitempty" xml:"order_detail_fee,omitempty"`
}

var poolEleOrderInfoBo = sync.Pool{
	New: func() any {
		return new(EleOrderInfoBo)
	},
}

// GetEleOrderInfoBo() 从对象池中获取EleOrderInfoBo
func GetEleOrderInfoBo() *EleOrderInfoBo {
	return poolEleOrderInfoBo.Get().(*EleOrderInfoBo)
}

// ReleaseEleOrderInfoBo 释放EleOrderInfoBo
func ReleaseEleOrderInfoBo(v *EleOrderInfoBo) {
	v.ResponsibleParty = ""
	v.Amount = ""
	v.TradeCreateTime = ""
	v.OrderCreateTime = ""
	v.PayEntity = ""
	v.EleOrderId = ""
	v.OrderId = ""
	v.OrderFrom = ""
	v.OrderIndex = ""
	v.OrderDetailFee = nil
	poolEleOrderInfoBo.Put(v)
}
