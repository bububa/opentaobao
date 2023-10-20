package happytrip

import (
	"sync"
)

// PassengerOrderInfo 结构体
type PassengerOrderInfo struct {
	// 订单id
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 订单状态码
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 订单详细状态码
	SubStatus string `json:"sub_status,omitempty" xml:"sub_status,omitempty"`
	// 出发地经度
	Flng string `json:"flng,omitempty" xml:"flng,omitempty"`
	// 出发地纬度
	Flat string `json:"flat,omitempty" xml:"flat,omitempty"`
	// 目的地经度
	Tlng string `json:"tlng,omitempty" xml:"tlng,omitempty"`
	// 目的地纬度
	Tlat string `json:"tlat,omitempty" xml:"tlat,omitempty"`
	// 乘客人数
	PassengerNumber int64 `json:"passenger_number,omitempty" xml:"passenger_number,omitempty"`
}

var poolPassengerOrderInfo = sync.Pool{
	New: func() any {
		return new(PassengerOrderInfo)
	},
}

// GetPassengerOrderInfo() 从对象池中获取PassengerOrderInfo
func GetPassengerOrderInfo() *PassengerOrderInfo {
	return poolPassengerOrderInfo.Get().(*PassengerOrderInfo)
}

// ReleasePassengerOrderInfo 释放PassengerOrderInfo
func ReleasePassengerOrderInfo(v *PassengerOrderInfo) {
	v.Id = ""
	v.Status = ""
	v.SubStatus = ""
	v.Flng = ""
	v.Flat = ""
	v.Tlng = ""
	v.Tlat = ""
	v.PassengerNumber = 0
	poolPassengerOrderInfo.Put(v)
}
