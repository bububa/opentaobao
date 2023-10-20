package wdk

import (
	"sync"
)

// EleBillBo 结构体
type EleBillBo struct {
	// 订单列表
	OrderList []EleOrderInfoBo `json:"order_list,omitempty" xml:"order_list>ele_order_info_bo,omitempty"`
	// 账单日期，时间戳
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 渠道店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 应付金额
	ExpendFee string `json:"expend_fee,omitempty" xml:"expend_fee,omitempty"`
	// 单量
	OrderCount string `json:"order_count,omitempty" xml:"order_count,omitempty"`
	// 未结算金额，单位:分
	PayFee string `json:"pay_fee,omitempty" xml:"pay_fee,omitempty"`
	// 实际付款主体
	PayEntity string `json:"pay_entity,omitempty" xml:"pay_entity,omitempty"`
	// 订单费用明细
	OrderDetailFee *OrderDetailFee `json:"order_detail_fee,omitempty" xml:"order_detail_fee,omitempty"`
}

var poolEleBillBo = sync.Pool{
	New: func() any {
		return new(EleBillBo)
	},
}

// GetEleBillBo() 从对象池中获取EleBillBo
func GetEleBillBo() *EleBillBo {
	return poolEleBillBo.Get().(*EleBillBo)
}

// ReleaseEleBillBo 释放EleBillBo
func ReleaseEleBillBo(v *EleBillBo) {
	v.OrderList = v.OrderList[:0]
	v.Date = ""
	v.ShopId = ""
	v.ExpendFee = ""
	v.OrderCount = ""
	v.PayFee = ""
	v.PayEntity = ""
	v.OrderDetailFee = nil
	poolEleBillBo.Put(v)
}
