package tblogistics

import (
	"sync"
)

// CarriageDetail 结构体
type CarriageDetail struct {
	// 破损赔付
	DamagePayment string `json:"damage_payment,omitempty" xml:"damage_payment,omitempty"`
	// 物流公司揽收时间段
	GotTime string `json:"got_time,omitempty" xml:"got_time,omitempty"`
	// 丢单赔付
	LostPayment string `json:"lost_payment,omitempty" xml:"lost_payment,omitempty"`
	// 快件送达所需的时间(单位：天)
	WayDay string `json:"way_day,omitempty" xml:"way_day,omitempty"`
	// 续费（单位：元）
	AddFee int64 `json:"add_fee,omitempty" xml:"add_fee,omitempty"`
	// 续重（单位：千克）
	AddWeight int64 `json:"add_weight,omitempty" xml:"add_weight,omitempty"`
	// 首费（单位：元）
	InitialFee int64 `json:"initial_fee,omitempty" xml:"initial_fee,omitempty"`
	// 首重（单位：千克）
	InitialWeight int64 `json:"initial_weight,omitempty" xml:"initial_weight,omitempty"`
}

var poolCarriageDetail = sync.Pool{
	New: func() any {
		return new(CarriageDetail)
	},
}

// GetCarriageDetail() 从对象池中获取CarriageDetail
func GetCarriageDetail() *CarriageDetail {
	return poolCarriageDetail.Get().(*CarriageDetail)
}

// ReleaseCarriageDetail 释放CarriageDetail
func ReleaseCarriageDetail(v *CarriageDetail) {
	v.DamagePayment = ""
	v.GotTime = ""
	v.LostPayment = ""
	v.WayDay = ""
	v.AddFee = 0
	v.AddWeight = 0
	v.InitialFee = 0
	v.InitialWeight = 0
	poolCarriageDetail.Put(v)
}
