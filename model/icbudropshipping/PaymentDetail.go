package icbudropshipping

import (
	"sync"
)

// PaymentDetail 结构体
type PaymentDetail struct {
	// shipment fee
	ShipmentFee string `json:"shipment_fee,omitempty" xml:"shipment_fee,omitempty"`
	// total amount
	TotalAmount string `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
}

var poolPaymentDetail = sync.Pool{
	New: func() any {
		return new(PaymentDetail)
	},
}

// GetPaymentDetail() 从对象池中获取PaymentDetail
func GetPaymentDetail() *PaymentDetail {
	return poolPaymentDetail.Get().(*PaymentDetail)
}

// ReleasePaymentDetail 释放PaymentDetail
func ReleasePaymentDetail(v *PaymentDetail) {
	v.ShipmentFee = ""
	v.TotalAmount = ""
	poolPaymentDetail.Put(v)
}
