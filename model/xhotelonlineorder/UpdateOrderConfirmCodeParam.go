package xhotelonlineorder

import (
	"sync"
)

// UpdateOrderConfirmCodeParam 结构体
type UpdateOrderConfirmCodeParam struct {
	// PMS确认号
	PmsResId string `json:"pms_res_id,omitempty" xml:"pms_res_id,omitempty"`
	// 商家系统订单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 飞猪订单号
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
}

var poolUpdateOrderConfirmCodeParam = sync.Pool{
	New: func() any {
		return new(UpdateOrderConfirmCodeParam)
	},
}

// GetUpdateOrderConfirmCodeParam() 从对象池中获取UpdateOrderConfirmCodeParam
func GetUpdateOrderConfirmCodeParam() *UpdateOrderConfirmCodeParam {
	return poolUpdateOrderConfirmCodeParam.Get().(*UpdateOrderConfirmCodeParam)
}

// ReleaseUpdateOrderConfirmCodeParam 释放UpdateOrderConfirmCodeParam
func ReleaseUpdateOrderConfirmCodeParam(v *UpdateOrderConfirmCodeParam) {
	v.PmsResId = ""
	v.OutOrderId = ""
	v.Tid = 0
	poolUpdateOrderConfirmCodeParam.Put(v)
}
