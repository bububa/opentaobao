package tmallhk

import (
	"sync"
)

// OrderCertify 结构体
type OrderCertify struct {
	// 订购人身份证号加密而成的加密串
	OcrId string `json:"ocr_id,omitempty" xml:"ocr_id,omitempty"`
	// 订购人姓名加密而成的加密串
	OcrName string `json:"ocr_name,omitempty" xml:"ocr_name,omitempty"`
	// 身份证的幂等号，用来标识订单实名的唯一性
	Idempotent string `json:"idempotent,omitempty" xml:"idempotent,omitempty"`
	// 订购人电话加密而成的加密串
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 订单编号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolOrderCertify = sync.Pool{
	New: func() any {
		return new(OrderCertify)
	},
}

// GetOrderCertify() 从对象池中获取OrderCertify
func GetOrderCertify() *OrderCertify {
	return poolOrderCertify.Get().(*OrderCertify)
}

// ReleaseOrderCertify 释放OrderCertify
func ReleaseOrderCertify(v *OrderCertify) {
	v.OcrId = ""
	v.OcrName = ""
	v.Idempotent = ""
	v.Phone = ""
	v.OrderId = 0
	poolOrderCertify.Put(v)
}
