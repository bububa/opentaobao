package examination

import (
	"sync"
)

// OrderRefundVo 结构体
type OrderRefundVo struct {
	// 子订单信息
	SubOrderList []SubOrderRequest `json:"sub_order_list,omitempty" xml:"sub_order_list>sub_order_request,omitempty"`
	// 主订单
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 退款id
	RefundOrderId string `json:"refund_order_id,omitempty" xml:"refund_order_id,omitempty"`
}

var poolOrderRefundVo = sync.Pool{
	New: func() any {
		return new(OrderRefundVo)
	},
}

// GetOrderRefundVo() 从对象池中获取OrderRefundVo
func GetOrderRefundVo() *OrderRefundVo {
	return poolOrderRefundVo.Get().(*OrderRefundVo)
}

// ReleaseOrderRefundVo 释放OrderRefundVo
func ReleaseOrderRefundVo(v *OrderRefundVo) {
	v.SubOrderList = v.SubOrderList[:0]
	v.OrderId = ""
	v.RefundOrderId = ""
	poolOrderRefundVo.Put(v)
}
