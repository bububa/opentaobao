package flight

import (
	"sync"
)

// VirProOrderVo 结构体
type VirProOrderVo struct {
	// 乘机人购买辅营产品详情
	PassengerAuxVos []PassengerAuxVo `json:"passenger_aux_vos,omitempty" xml:"passenger_aux_vos>passenger_aux_vo,omitempty"`
	// 预订时间，辅营订单创建时间
	BookTime string `json:"book_time,omitempty" xml:"book_time,omitempty"`
	// 支付宝流水号，存在多个辅营订单对应一个支付宝流水号的情况
	PayNo string `json:"pay_no,omitempty" xml:"pay_no,omitempty"`
	// 支付时间，订单为支付成功或出货成功时返回
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 联系人电话
	ContactPhone string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// 机票订单号
	FlightOrderId int64 `json:"flight_order_id,omitempty" xml:"flight_order_id,omitempty"`
	// 辅营订单金额
	PayPrice int64 `json:"pay_price,omitempty" xml:"pay_price,omitempty"`
	// 辅营订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 辅营订单状态，1-	待支付 2-	支付成功 3-	辅营出货成功 4-	订单取消
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
}

var poolVirProOrderVo = sync.Pool{
	New: func() any {
		return new(VirProOrderVo)
	},
}

// GetVirProOrderVo() 从对象池中获取VirProOrderVo
func GetVirProOrderVo() *VirProOrderVo {
	return poolVirProOrderVo.Get().(*VirProOrderVo)
}

// ReleaseVirProOrderVo 释放VirProOrderVo
func ReleaseVirProOrderVo(v *VirProOrderVo) {
	v.PassengerAuxVos = v.PassengerAuxVos[:0]
	v.BookTime = ""
	v.PayNo = ""
	v.PayTime = ""
	v.ContactPhone = ""
	v.FlightOrderId = 0
	v.PayPrice = 0
	v.OrderId = 0
	v.OrderStatus = 0
	poolVirProOrderVo.Put(v)
}
