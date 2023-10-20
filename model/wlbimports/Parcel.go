package wlbimports

import (
	"sync"
)

// Parcel 结构体
type Parcel struct {
	// 小包LP单号
	LgOrderCode string `json:"lg_order_code,omitempty" xml:"lg_order_code,omitempty"`
	// 包裹运单号
	TrackingNumber string `json:"tracking_number,omitempty" xml:"tracking_number,omitempty"`
	// 0: 成功 异常详见： 2：少件 3：多件 5：其他
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
}

var poolParcel = sync.Pool{
	New: func() any {
		return new(Parcel)
	},
}

// GetParcel() 从对象池中获取Parcel
func GetParcel() *Parcel {
	return poolParcel.Get().(*Parcel)
}

// ReleaseParcel 释放Parcel
func ReleaseParcel(v *Parcel) {
	v.LgOrderCode = ""
	v.TrackingNumber = ""
	v.Status = ""
	v.Remark = ""
	poolParcel.Put(v)
}
